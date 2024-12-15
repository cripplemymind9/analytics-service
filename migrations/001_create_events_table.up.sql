-- Создание таблицы events
CREATE TABLE events (
    user_id String,
    url String,
    timestamp DateTime
) ENGINE = MergeTree()
ORDER BY (user_id, timestamp)
PARTITION BY toYYYYMM(timestamp)
SETTINGS index_granularity = 8192;
