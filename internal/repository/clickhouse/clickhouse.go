package clickhouse

import (
	"go.uber.org/zap"

	impl "github.com/cripplemymind9/analytics-service/internal/repository/clickhouse/impl"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/clickhouse/clickhouse_event"
	"github.com/cripplemymind9/analytics-service/internal/ep/clickhouse"
)

type Repositories struct {
	clickhouse_event.Event
}

func NewRepositories(ch *clickhouse.ClickhouseClient, logger *zap.Logger) *Repositories {
	return &Repositories{
		Event: impl.NewEventRepo(ch, logger),
	}
}