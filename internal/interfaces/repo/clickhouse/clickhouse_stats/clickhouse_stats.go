package clickhouse_stats

import (
	"context"

	"github.com/cripplemymind9/analytics-service/internal/models"
)

type Stats interface {
	GetStats(ctx context.Context, fromTime, toTime string) (*models.StatsData, error)
}
