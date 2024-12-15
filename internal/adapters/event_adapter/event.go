package event_adapter

import (
	"context"
	"fmt"
	"time"

	"github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

func (s *Service) AddEvent(ctx context.Context, req *event.AddEventRequest) (*event.AddEventResponse, error) {
	if req.UserId == "" || req.Url == "" || req.Timestamp == "" {
		return &event.AddEventResponse{
			Success: false,
			Message: "All fields are required",
		}, nil
	}

	// Проверка формата временной метки
	_, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return &event.AddEventResponse{
			Success: false,
			Message: "Invalid timestamp format. Use ISO8601 (RFC3339)",
		}, nil
	}

	if err := s.svc.AddEvent(ctx, req); err != nil {
		return &event.AddEventResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to insert event: %v", err),
		}, nil
	}

	return &event.AddEventResponse{
		Success: true,
		Message: "Event added successfully",
	}, nil
}