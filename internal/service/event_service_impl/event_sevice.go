package event_service_impl

import (
	"context"
	"time"
	"fmt"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/clickhouse/clickhouse_event"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/cache/cache_event"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/event_service"
	"github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

var _ event_service.Event = (*EventService)(nil)

type EventService struct {
	eventRepo clickhouse_event.Event
	eventCache cache_event.Event
}

func New(eventRepo clickhouse_event.Event, eventCache cache_event.Event) event_service.Event {
	return &EventService{
		eventRepo: eventRepo,
		eventCache: eventCache,
	}
}

func (s *EventService) AddEvent(ctx context.Context, req *event.AddEventRequest) error {
	parsedTime, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}
	formattedTime := parsedTime.Format("2006-01-02 15:04:05")

	err = s.eventRepo.AddEvent(ctx, req.UserId, req.Url, formattedTime)
	if err != nil {
		return fmt.Errorf("add event: %w", err)
	}

	err = s.eventCache.AddEvent(ctx, req.UserId, req.Url, req.Timestamp)
	if err != nil {
		return fmt.Errorf("add event: %w", err)
	}

	return nil
}