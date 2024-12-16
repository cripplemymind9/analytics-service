package stats_service_impl

import (
	"context"
	"fmt"
	"time"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/repo/clickhouse/clickhouse_stats"
	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/stats_service"
	
	"github.com/cripplemymind9/analytics-service/pkg/pb/stats"
)

var _ stats_service.Stats = (*StatsService)(nil)

type StatsService struct {
	statsRepo clickhouse_stats.Stats
}

func New(statsRepo clickhouse_stats.Stats) stats_service.Stats {
	return &StatsService{
		statsRepo: statsRepo,
	}
}

func (s *StatsService) GetStats(
	ctx context.Context,
	req *stats.GetStatsRequest,
	) (*stats.GetStatsResponse, error) {

	formattedFromTime, err := parseAndFormatTime(req.From)
	if err != nil {
		return nil, fmt.Errorf("invalid 'from' parameter: %w", err)
	}
	
	formattedToTime, err := parseAndFormatTime(req.To)
	if err != nil {
		return nil, fmt.Errorf("invalid 'to' parameter: %w", err)
	}

	// Полученеи данных из репоизтория.
	statsData, err := s.statsRepo.GetStats(ctx, formattedFromTime, formattedToTime)
	if err != nil {
		return nil, fmt.Errorf("repository get stats: %w", err)
	}

	// Формирование ответа.
	response := &stats.GetStatsResponse{
		UniqueUsers: int32(statsData.UniqueUsers),
		TotalEvents: int32(statsData.TotalEvents),
		MostVisitedUrls: make([]*stats.MostVisitedUrl, len(statsData.MostVisitedUrls)),
	}

	for i, url := range statsData.MostVisitedUrls {
		response.MostVisitedUrls[i] = &stats.MostVisitedUrl{
			Url: url.Url,
			Count: int32(url.Count),
		}
	}

	return response, nil
}

// parseAndFormatTime- функция проверки формата времени
func parseAndFormatTime(inputTime string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		return "", err
	}
	return parsedTime.Format("2006-01-02 15:04:05"), nil
}
