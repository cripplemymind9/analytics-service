package ep

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep/config"
	"github.com/cripplemymind9/analytics-service/internal/ep/clickhouse"
	"github.com/cripplemymind9/analytics-service/internal/ep/redis"
	"github.com/cripplemymind9/analytics-service/internal/server"
	"github.com/cripplemymind9/analytics-service/internal/adapters/health_adapter"
)

// Run запускает основной процесс сервера, включая инициализацию всех зависимостей
func Run(ctx context.Context, cfg *config.Config, logger *zap.Logger) error {
	clickhouseClient, err := clickhouse.NewClient(cfg, logger)
	if err != nil {
		logger.Error("Failed to create ClickHouse client", zap.Error(err))
		return fmt.Errorf("create clickhouseClient: %w", err)
	}
	defer clickhouseClient.Close()
	logger.Info("ClickHouse client created successfully", zap.String("address", cfg.ClickHouse.Port))

	redisClient, err := redis.NewClient(ctx, cfg, cfg.Redis.DB, logger)
	if err != nil {
		logger.Error("Failed to create Redis client", zap.Error(err))
		return fmt.Errorf("create redisClient: %w", err)
	}
	defer redisClient.Close()
	logger.Info("Redis client created successfully", zap.String("address", cfg.Redis.Port))

	server, err := server.New(
		cfg,
		logger,
		server.WithImplementationAdapters(
			health_adapter.New(),
		),
	)
	if err != nil {
		return fmt.Errorf("create server: %w", err)
	}
	logger.Info("Starting gRPC server", zap.String("gRPC port", cfg.GRPC.Port))
	
	quit := setupSignalChannel()
	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- server.Start(ctx)
	}()

	// Ожидание ошибки сервера или сигнала завершения
	select {
	case err = <-serverErrors:
		return fmt.Errorf("gRPC server encountered an error: %w", err)
	case sig := <-quit:
		logger.Warn("Termination signal received", zap.String("signal", sig.String()))
	}

	// Корректное завершение работы gRPC сервера
	logger.Info("Shutting down gRPC server gracefully...")
	server.Stop()
	logger.Info("gRPC server stopped")
	return nil
}

// setupSignalChannel - настраивает канал для прослушивания системных сигналов завершения
func setupSignalChannel() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	return quit
}