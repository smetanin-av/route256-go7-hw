package main

//goland:noinspection GoSnakeCaseUsage
import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"route256/libs/close_funcs"
	"route256/libs/interceptors/panic_recover"
	"route256/libs/interceptors/validating"
	kafka_producer "route256/libs/kafka/producer"
	"route256/libs/logging"
	"route256/libs/tracing"
	"route256/libs/tx_manager"
	"route256/loms/internal/app/config"
	loms_repo "route256/loms/internal/app/repository/loms"
	kafka_sender "route256/loms/internal/app/sender/kafka"
	loms_server "route256/loms/internal/app/server/loms"
	loms_service "route256/loms/internal/app/service/loms"
	"route256/loms/pkg/loms/v1"
)

const (
	shutdownTimeout = 5 * time.Second
)

var (
	closer = close_funcs.New()
)

func init() {
	err := logging.InitGlobal()
	if err != nil {
		panic(err)
	}

	err = tracing.InitGlobal("loms")
	if err != nil {
		panic(err)
	}
}

func main() {
	bCtx := context.Background()
	ctx, cancel := signal.NotifyContext(bCtx, syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		cancel()
		shutdownApp(bCtx)
	}()

	logging.Infof("starting app")
	if err := runApp(ctx); err != nil {
		logging.Errorf("run app: %v", err)
	}
}

func runApp(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	serverImpl, err := initServerImpl(ctx, cfg)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GrpcPort))
	if err != nil {
		return fmt.Errorf("net listen: %w", err)
	}

	errs := make(chan error, 1)
	go func() {
		errs <- runGrpcServer(serverImpl, listener)
	}()
	go func() {
		errs <- runHttpGateway(ctx, listener.Addr().String(), cfg.HttpPort)
	}()

	select {
	case <-ctx.Done():
		return nil
	case err = <-errs:
		return err
	}
}

func shutdownApp(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	logging.Infof("shutting down app gracefully")
	if err := closer.Close(ctx); err != nil {
		logging.Errorf("shutdown app: %v", err)
	}
}

func initServerImpl(ctx context.Context, cfg *config.Config) (*loms_server.Server, error) {
	pool, err := pgxpool.Connect(ctx, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}
	closer.Add(pool.Close)

	provider := tx_manager.New(pool)
	repo := loms_repo.New(provider)

	producer, err := kafka_producer.New(cfg.KafkaBrokers)
	if err != nil {
		return nil, fmt.Errorf("init kafka producer: %w", err)
	}
	closer.Add(producer.Close)

	sender := kafka_sender.New(producer, cfg.KafkaTopic)
	service := loms_service.New(provider, repo, sender)

	go runCleanUpOrders(ctx, service, cfg.MaxAwaitingPayment)

	return loms_server.New(service), nil
}

func runGrpcServer(serverImpl *loms_server.Server, listener net.Listener) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(panic_recover.Interceptor),
		grpc.ChainUnaryInterceptor(logging.Interceptor),
		grpc.ChainUnaryInterceptor(tracing.Interceptor),
		grpc.ChainUnaryInterceptor(validating.Interceptor),
	)

	reflection.Register(grpcServer)

	loms_v1.RegisterLomsServiceServer(grpcServer, serverImpl)

	logging.Infof("grpc listening on %s", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("grpc serve: %v", err)
	}

	return nil
}

func runHttpGateway(ctx context.Context, grpcHost string, port uint32) error {
	conn, err := grpc.DialContext(
		ctx,
		grpcHost,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("grpc dial: %w", err)
	}

	mux := runtime.NewServeMux()
	err = loms_v1.RegisterLomsServiceHandler(ctx, mux, conn)
	if err != nil {
		return fmt.Errorf("register gateway: %w", err)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	logging.Infof("http listening on %d", port)
	err = httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %v", err)
	}

	return nil
}

func runCleanUpOrders(ctx context.Context, service *loms_service.Service, delta time.Duration) {
	// случайная задержка, чтобы реплики сервиса не штурмовали БД все вместе
	delay := rand.Intn(30) + 30
	time.Sleep(time.Second * time.Duration(delay))

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			service.CleanUpOrders(ctx, delta)

		default:
			time.Sleep(time.Second * 15)
		}
	}
}
