package service

import (
	"github.com/cripplemymind9/analytics-service/internal/repository/clickhouse"
	"github.com/cripplemymind9/analytics-service/internal/repository/cache_impl"

	event_service_impl "github.com/cripplemymind9/analytics-service/internal/service/event_service_impl"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/event_service"
)

type Services struct {
	Event event_service.Event
	Repos clickhouse.Repositories
	Cache cache_impl.Cache
}

func New(repos clickhouse.Repositories, cache cache_impl.Cache) *Services {
	return &Services{
		Event: event_service_impl.New(repos, &cache),
	}
}