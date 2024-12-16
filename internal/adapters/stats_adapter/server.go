package stats_adapter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/stats_service"
	pb "github.com/cripplemymind9/analytics-service/pkg/pb/stats"
)

type Service struct {
	pb.UnimplementedStatsServiceServer
	svc stats_service.Stats
}

func New(svc stats_service.Stats) *Service {
	return &Service{svc: svc}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterStatsServiceServer(server, s)
}

func (s *Service) RegisterHandler(
	ctx context.Context,
	mux *runtime.ServeMux,
	conn *grpc.ClientConn,
	) error {
	return pb.RegisterStatsServiceHandler(ctx, mux, conn)
}
