package clickhouse

import (
	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep/clickhouse"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/clickhouse/clickhouse_event"
	impl "github.com/cripplemymind9/analytics-service/internal/repository/clickhouse/impl"
)

type Repositories struct {
	clickhouse_event.Event
}

func NewRepositories(ch *clickhouse.ClickhouseClient, logger *zap.Logger) *Repositories {
	return &Repositories{
		Event: impl.NewEventRepo(ch, logger),
	}
}
