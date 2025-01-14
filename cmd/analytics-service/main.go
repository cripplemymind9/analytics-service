package main

import (
	"context"
	"log"

	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep"
	"github.com/cripplemymind9/analytics-service/internal/ep/config"
	"github.com/cripplemymind9/analytics-service/internal/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	logger := logger.NewClientZapLogger(cfg.Log.Level, cfg.DefaultClientId)

	if err = ep.Run(ctx, cfg, logger); err != nil {
		logger.Fatal("Run server failed: %s\n", zap.Error(err))
	}
}
