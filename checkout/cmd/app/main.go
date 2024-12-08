package main

//goland:noinspection GoSnakeCaseUsage
import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	loms_client "route256/checkout/internal/app/client/loms"
	product_client "route256/checkout/internal/app/client/product"
	"route256/checkout/internal/app/config"
	checkout_repo "route256/checkout/internal/app/repository/checkout"
	checkout_server "route256/checkout/internal/app/server/checkout"
	checkout_service "route256/checkout/internal/app/service/checkout"
	"route256/checkout/pkg/checkout/v1"
	"route256/libs/close_funcs"
	"route256/libs/interceptors/panic_recover"
	"route256/libs/interceptors/validating"
	"route256/libs/logging"
	"route256/libs/tracing"
	"route256/libs/tx_manager"
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

	err = tracing.InitGlobal("checkout")
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

func initServerImpl(ctx context.Context, cfg *config.Config) (*checkout_server.Server, error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: "",
		DB:       0,
	})

	lomsCli, err := loms_client.New(ctx, cfg.LomsService)
	if err != nil {
		return nil, fmt.Errorf("init loms client: %w", err)
	}
	closer.Add(lomsCli.Close)

	productCli, err := product_client.New(ctx, cfg.ProductService, cfg.Token, redisCli)
	if err != nil {
		return nil, fmt.Errorf("init product client: %w", err)
	}
	closer.Add(productCli.Close)

	pool, err := pgxpool.Connect(ctx, cfg.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}
	closer.Add(pool.Close)

	provider := tx_manager.New(pool)
	service := checkout_service.New(lomsCli, productCli, provider, checkout_repo.New(provider))

	return checkout_server.New(service), nil
}

func runGrpcServer(serverImpl *checkout_server.Server, listener net.Listener) error {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(panic_recover.Interceptor),
		grpc.ChainUnaryInterceptor(logging.Interceptor),
		grpc.ChainUnaryInterceptor(tracing.Interceptor),
		grpc.ChainUnaryInterceptor(validating.Interceptor),
	)

	reflection.Register(grpcServer)

	checkout_v1.RegisterCheckoutServiceServer(grpcServer, serverImpl)

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
	err = checkout_v1.RegisterCheckoutServiceHandler(ctx, mux, conn)
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
