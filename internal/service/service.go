package service

import (
	event_service_impl "github.com/cripplemymind9/analytics-service/internal/service/event_service_impl"
	stats_service_impl "github.com/cripplemymind9/analytics-service/internal/service/stats_service_impl"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/event_service"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/stats_service"

	"github.com/cripplemymind9/analytics-service/internal/repository/clickhouse"

)

type Services struct {
	Event event_service.Event
	Stats stats_service.Stats
	Repos clickhouse.Repositories
}

func New(repos clickhouse.Repositories) *Services {
	return &Services{
		Event: event_service_impl.New(repos),
		Stats: stats_service_impl.New(repos),
	}
}
