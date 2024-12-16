package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
)

// Структура для хранения конфигурации приложения.
type (
	Config struct {
		// Общие параметры приложения.
		App struct {
			Name string `envconfig:"APP_NAME" required:"true"`
			Version string 	`envconfig:"APP_VERSION" required:"true"`
		}

		// Параметры для GRPC сервера.
		GRPC struct {
			Port string `envconfig:"GRPC_PORT" required:"true"`
		}

		// Параметры для GW.
		Gateway struct {
			Port string `envconfig:"GW_PORT" required:"true"`
		}

		// Уровень логирования.
		Log struct {
			Level string `envconfig:"LOG_LEVEL" required:"true"`
		}

		// Параметры для Clickhouse.
		ClickHouse struct {
			Host     string `envconfig:"CLICKHOUSE_HOST" required:"true"`
			Port     string `envconfig:"CLICKHOUSE_PORT" required:"true"`
			Database string `envconfig:"CLICKHOUSE_DB" default:"default"`
		}

		// Параметры среды выполнения и ID клиента.
		DefaultClientId string 	`envconfig:"DEFAULT_CLIENT_ID" required:"true"`
		Env 			string  `envconfig:"ENV" default:"dev"` 	
	}
)

// GetConfigFromEnv загружает конфигурации из .env файла и переменных окружения.
func GetConfigFromEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load the .env file: %s\n", err.Error())
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
