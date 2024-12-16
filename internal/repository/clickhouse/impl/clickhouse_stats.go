package impl

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep/clickhouse"
	"github.com/cripplemymind9/analytics-service/internal/models"
)

type StatsRepo struct {
	ch 		*clickhouse.ClickhouseClient
	logger 	*zap.Logger
}

func NewStatsRepo(ch *clickhouse.ClickhouseClient, logger *zap.Logger) *StatsRepo {
	return &StatsRepo{
		ch,
		logger,
	}
}

func (r *StatsRepo) GetStats(
	ctx context.Context,
	fromTime,
	toTime string,
	) (*models.StatsData, error) {
	// Запрос на получение уникальных пользователей и количества событий
	statsQuery := `
		SELECT 
			COUNT(DISTINCT user_id) AS unique_users, 
			COUNT(*) AS total_events
		FROM events
		WHERE timestamp BETWEEN ? AND ?
	`

	var uniqueUsers, totalEvents int
	err := r.ch.Db.QueryRowContext(ctx, statsQuery, fromTime, toTime).Scan(&uniqueUsers, &totalEvents)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats summary: %w", err)
	}

	// Запрос на получение самых посещаемых URL
	urlsQuery := `
		SELECT 
			url, 
			COUNT(*) AS visit_count
		FROM events
		WHERE timestamp BETWEEN ? AND ?
		GROUP BY url
		ORDER BY visit_count DESC
		LIMIT 10
	`
	rows, err := r.ch.Db.QueryContext(ctx, urlsQuery, fromTime, toTime)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch most visited URLs: %w", err)
	}
	defer rows.Close()

	var mostVisitedUrls []models.MostVisitedUrlData
	for rows.Next() {
		var urlData models.MostVisitedUrlData
		err := rows.Scan(&urlData.Url, &urlData.Count)
		if err != nil {
			return nil, fmt.Errorf("failed to scan URL data: %w", err)
		}
		mostVisitedUrls = append(mostVisitedUrls, urlData)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	statsData := &models.StatsData{
		UniqueUsers: uniqueUsers,
		TotalEvents: totalEvents,
		MostVisitedUrls: mostVisitedUrls,
	}

	return statsData, nil
}