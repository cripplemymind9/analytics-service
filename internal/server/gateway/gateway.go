package gateway

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/cripplemymind9/analytics-service/internal/server/adapter"
)

// Структура HTTP Gateway.
type Gateway struct {
	logger     *zap.Logger
	adapters   []adapter.ImplementationAdapter
	gatewayMux *http.ServeMux
}

// New - создает новый экземпляр Gateway.
func New(logger *zap.Logger, adapters []adapter.ImplementationAdapter) *Gateway {
	return &Gateway{
		logger:     logger,
		adapters:   adapters,
		gatewayMux: http.NewServeMux(),
	}
}

// Start - запускает HTTP Gateway для gRPC сервера.
func (g *Gateway) Start(ctx context.Context, httpAddress, grpcAddress string) error {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		g.logger.Error("failed to connect to gRPC server", zap.Error(err))
		return err
	}
	defer conn.Close()

	// Настраиваем JSON-маршаллер
	defaultMarshallerOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   false,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	opts := []runtime.ServeMuxOption{
		defaultMarshallerOption,
	}

	gwmux := runtime.NewServeMux(opts...)

	// Решистрируем обработчики из списка адаптеров
	for _, impl := range g.adapters {
		if err := impl.RegisterHandler(ctx, gwmux, conn); err != nil {
			g.logger.Error("failed to register handler", zap.Error(err))
			return err
		}
	}

	// Создаем HTTP-mux и добавляем Gateway
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	httpServer := &http.Server{
		Addr:              httpAddress,
		Handler:           mux,
		ReadHeaderTimeout: 15 * time.Second,
	}

	g.logger.Info("gRPC Gateway starting", zap.String("address", httpAddress))

	// Запускаем HTTP-сервер
	if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		g.logger.Error("failed to start http gateway server", zap.Error(err))
		return err
	}

	return nil
}
