package impl

import (
	"context"

	"go.uber.org/zap"

	"github.com/cripplemymind9/analytics-service/internal/ep/clickhouse"
)

type EventRepo struct {
	ch 		*clickhouse.ClickhouseClient
	logger 	*zap.Logger
}

func NewEventRepo(ch *clickhouse.ClickhouseClient, logger *zap.Logger) *EventRepo {
	return &EventRepo{
		ch,
		logger,
	}
}

func (r *EventRepo) AddEvent(
		ctx context.Context,
		userId,
		url,
		timestamp string,
	) error {
	tx, err := r.ch.Db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to begin transaction", zap.Error(err))
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO events (user_id, url, timestamp) VALUES (?, ?, ?)
	`

	_, err = tx.ExecContext(ctx, query, userId, url, timestamp)
	if err != nil {
		r.logger.Error("failed to add event", zap.Error(err))
		return err
	}

	if err = tx.Commit(); err != nil {
		r.logger.Error("failed to commit transaction", zap.Error(err))
		return err
	}

	r.logger.Info("event added successfully", zap.String("user_id", userId), zap.String("url", url))
	return nil
}
