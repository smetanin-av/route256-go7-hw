package main

//goland:noinspection GoSnakeCaseUsage
import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"route256/libs/close_funcs"
	kafka_consumer "route256/libs/kafka/consumer"
	"route256/libs/tx_manager"
	"route256/notifications/internal/app/config"
	"route256/notifications/internal/app/notifier/logger"
	"route256/notifications/internal/app/notifier/telegram"
	kafka_receiver "route256/notifications/internal/app/receiver/kafka"
	"route256/notifications/internal/app/repository"
	"route256/notifications/internal/app/service"
)

const (
	shutdownTimeout = 5 * time.Second
)

var (
	closer = close_funcs.New()
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		cancel()
		shutdownApp()
	}()

	if err := runApp(ctx); err != nil {
		log.Println("ERROR run app:", err)
	}
}

func runApp(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pool, err := pgxpool.Connect(ctx, cfg.DatabaseDSN)
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	closer.Add(pool.Close)

	provider := tx_manager.New(pool)
	repo := repository.New(provider)

	consumer, err := kafka_consumer.New(cfg.KafkaBrokers, cfg.ConsumerGroup)
	if err != nil {
		return fmt.Errorf("init kafka consumer: %w", err)
	}
	closer.Add(consumer.Close)

	var notifier service.Notifier
	if cfg.TgBotApiToken != "" && cfg.TgBotChatId != 0 {
		notifier, err = telegram.New(cfg.TgBotApiToken, cfg.TgBotChatId)
		if err != nil {
			return fmt.Errorf("init telegram api: %w", err)
		}
	} else {
		notifier = logger.New()
	}

	processor := service.New(provider, repo, notifier)
	kafka_receiver.Subscribe(ctx, consumer, cfg.KafkaTopic, processor.ProcessMsg)

	<-ctx.Done()
	return nil
}

func shutdownApp() {
	log.Println("shutting down server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := closer.Close(ctx); err != nil {
		log.Println("ERROR shutdown app:", err)
	}
}
