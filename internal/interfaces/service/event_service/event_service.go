package event_service

import (
	"context"
	"github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

type Event interface {
	AddEvent(ctx context.Context, req *event.AddEventRequest) error
}