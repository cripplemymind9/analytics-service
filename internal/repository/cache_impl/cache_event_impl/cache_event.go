package cache_event_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type EventCache struct {
	client 			*redis.Client
	logger 			*zap.Logger
	ttl 			time.Duration
	namespace 		string
}

func NewEventCache(client *redis.Client, timeTTL time.Duration, namespace string, logger *zap.Logger) *EventCache {
	return &EventCache{
		client: client,
		logger: logger,
		ttl: 	timeTTL,
		namespace: namespace,
	}
}

func (c *EventCache) AddEvent(ctx context.Context, userId, url, timestamp string) error {
	// Формирование ключа с учетом namesapce
	key := fmt.Sprintf("%s:event:%s:%s", c.namespace, userId, timestamp)

	// Значение — JSON-структура или хэш
	value := map[string]string{
		"userId":    userId,
		"url":       url,
		"timestamp": timestamp,
	}

	// Используем HSet для хранения данных в виде хэша
	err := c.client.HSet(ctx, key, value).Err()
	if err != nil {
		return fmt.Errorf("failed to cache event: %w", err)
	}

	// Устанавливаем время жизни ключа
	err = c.client.Expire(ctx, key, c.ttl).Err()
	if err != nil {
		return fmt.Errorf("failed to set TTL for event: %w", err)
	}

	return nil
}
