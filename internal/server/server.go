package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/cripplemymind9/analytics-service/internal/server/gateway"
	"github.com/cripplemymind9/analytics-service/internal/ep/config"
)

// Структура gRPC сервера с поддержкой HTTP-Gateway.
type Server struct {
	options
	logger      *zap.Logger
	grpcAddress string
	httpAddress string
	serviceName string
	grpcServer  *grpc.Server
}

// New - создает новый экземпляр gRPC-сервера.
func New(
	cfg *config.Config,
	logger *zap.Logger,
	opts ...EntrypointOption,
	) (*Server, error) {
	o := options{}

	for _, opt := range opts {
		opt.apply(&o)
	}

	return &Server{
		options: 			o,
		logger:				logger,
		grpcAddress: 		fmt.Sprintf("0.0.0.0:%s", cfg.GRPC.Port),
		httpAddress: 		fmt.Sprintf("0.0.0.0:%s", cfg.Gateway.Port),
		serviceName: 		cfg.App.Name,
	}, nil
}

// Start - запускает gRPC сервер и HTTP-Gateway.
func (s *Server) Start(ctx context.Context) error {
	// TODO: добавить интерсептор для сбора метрик (Prometheus)

	ints := []grpc.ServerOption{}

	if s.grpcUnaryServerInterceptors != nil {
		for _, v := range s.grpcUnaryServerInterceptors {
			ints = append(ints, grpc.ChainUnaryInterceptor(v))
		}
	}

	// Инициализация gRPC сервера
	s.grpcServer = grpc.NewServer(ints...)

	// TODO: инциализация метрик и запуск сервера Prometheus.

	// Настройка слушателя
	listener, err := net.Listen("tcp", s.grpcAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w", s.grpcAddress, err)
	}

	// Регистрация адаптеров
	for _, v := range s.adapters {
		v.RegisterServer(s.grpcServer)
		s.logger.Info("Registered gRPC handler", zap.String("adapter", fmt.Sprintf("%T", v)))
	}

	// Включаем поддержку ReflectionAPI для gRPC сервера.
	reflection.Register(s.grpcServer)

	// Запуск HTTP-Gateway
	gatewayServer := gateway.New(s.logger, s.adapters)
	go func() {
		if err := gatewayServer.Start(ctx, s.httpAddress, s.grpcAddress); err != nil {
			s.logger.Error("failed to start gRPC-Gateway", zap.Error(err))
			s.Stop()
		}
	}()

	// Запуск gRPC сервера.
	return s.grpcServer.Serve(listener)
}

// Stop - корректно завершает работу gRPC сервера.
func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
