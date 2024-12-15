package cache_impl

import (
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/repository/cache_impl/cache_event_impl"
)

const (
	timeTTL        = 3
	eventNameSpace = "event-service"
)

type Cache struct {
	cache_event_impl.EventCache
}

func New(client *redis.Client, logger *zap.Logger) *Cache {
	return &Cache{
		EventCache: *cache_event_impl.NewEventCache(client, timeTTL*time.Hour, eventNameSpace, logger),
	}
}
