FROM golang:1.22.7
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/analytics-service/main.go
# Устанавливаем dockerize
RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz && \
    tar -xvzf dockerize-linux-amd64-v0.6.1.tar.gz && \
    mv dockerize /usr/local/bin/

# Переменные окружения
ENV APP_NAME=analytics-service
ENV APP_VERSION=1.0.0

ENV GRPC_PORT=50051

ENV GW_PORT=8080

ENV LOG_LEVEL=debug

ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

ENV CLICKHOUSE_HOST=clickhouse
ENV CLICKHOUSE_PORT=9000

ENV DEFAULT_CLIENT_ID=my-client-id
ENV ENV=dev

CMD ["./main"]