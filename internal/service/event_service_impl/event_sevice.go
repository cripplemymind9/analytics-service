package event_service_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/clickhouse/clickhouse_event"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/event_service"
	
	"github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

var _ event_service.Event = (*EventService)(nil)

type EventService struct {
	eventRepo  clickhouse_event.Event
}

func New(eventRepo clickhouse_event.Event) event_service.Event {
	return &EventService{
		eventRepo:  eventRepo,
	}
}

func (s *EventService) AddEvent(
	ctx context.Context,
	req *event.AddEventRequest,
) error {
	// Проверка формата времени
	parsedTime, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}
	formattedTime := parsedTime.Format("2006-01-02 15:04:05")

	err = s.eventRepo.AddEvent(ctx, req.UserId, req.Url, formattedTime)
	if err != nil {
		return fmt.Errorf("add event: %w", err)
	}

	return nil
}
