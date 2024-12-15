package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep/config"
)

// RedisCLient - струкутра для работы с Redis.
type RedisClient struct {
	redisClient 	*redis.Client
	logger 			*zap.Logger
}

// New - создает новый экземпляр клиента Redis.
func NewClient(ctx context.Context, cfg *config.Config, db int, logger *zap.Logger) (*RedisClient, error) {
	url := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)

	// Создаем клиента Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: url,
		DB: db,
	})
	logger.Info("Redis client created", zap.String("address", url), zap.Int("db", db))

	// Проверяем подключение
	if err := redisClient.Ping(ctx).Err(); err != nil {
		logger.Error("Failed to ping Redis", zap.String("address", url), zap.Error(err))
		return nil, fmt.Errorf("could not connect to Redis at %s: %w", url, err)
	}
	logger.Info("Successfully connected to Redis")

	return &RedisClient{
		redisClient: 	redisClient,
		logger: 		logger,
	}, nil
}

// Close - закрытие соединения с Redis.
func (c *RedisClient) Close() error {
	c.logger.Info("Closing Redis connection")
	return c.redisClient.Close()
}