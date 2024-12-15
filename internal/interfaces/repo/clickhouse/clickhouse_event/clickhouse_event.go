package clickhouse_event

import "context"

type Event interface {
	AddEvent(ctx context.Context, userId, url, timestamp string) error
}