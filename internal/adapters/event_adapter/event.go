package event_adapter

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

func (s *Service) AddEvent(ctx context.Context, req *event.AddEventRequest) (*event.AddEventResponse, error) {
	if req.UserId == "" || req.Url == "" || req.Timestamp == "" {
		return nil, status.Error(codes.InvalidArgument, "All fields (user_id, url, timestamp) are required")
	}

	if err := s.svc.AddEvent(ctx, req); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert event: %v", err)
	}

	return &event.AddEventResponse{
		Success: true,
		Message: "Event added successfully",
	}, nil
}
