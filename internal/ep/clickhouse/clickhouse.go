package clickhouse

import (
	"database/sql"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/cripplemymind9/analytics-service/internal/ep/config"
	"go.uber.org/zap"
)

// ClickhouseClient - структура для работы с ClickHouse.
type ClickhouseClient struct {
	db 		*sql.DB
	logger 	*zap.Logger
}

// New - создаёт новый клиент для работы с ClickHouse
func NewClient(cfg *config.Config, logger *zap.Logger) (*ClickhouseClient, error) {
	dsn := fmt.Sprintf("clickhouse://default:@%s:%s?database=%s", cfg.ClickHouse.Host, cfg.ClickHouse.Port, cfg.ClickHouse.Database)
	
	// Подключение к базе данных
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		logger.Error("Failed to create ClickHouse client", zap.Error(err))
		return nil, fmt.Errorf("could not create ClickHouse client: %w", err)
	}

	// Проверка подключения
	if err := db.Ping(); err != nil {
		logger.Error("Failed to connect to ClickHouse", zap.String("dsn", dsn), zap.Error(err))
		return nil, fmt.Errorf("could not connect to ClickHouse: %w", err)
	}
	logger.Info("Successfully connected to ClickHouse", zap.String("dsn", dsn))

	return &ClickhouseClient{db: db, logger: logger}, nil
}

// Close - метод для закрытия соединения с ClickHouse.
func (c *ClickhouseClient) Close() error {
	c.logger.Info("Closing ClickHouse connection")
	if err := c.db.Close(); err != nil {
		c.logger.Error("Failed to close ClickHouse connection", zap.Error(err))
		return err
	}

	return nil
}