package stats_service

import (
	"context"
	"github.com/cripplemymind9/analytics-service/pkg/pb/stats"
)

type Stats interface {
	GetStats(ctx context.Context, req *stats.GetStatsRequest) (*stats.GetStatsResponse, error)
}
