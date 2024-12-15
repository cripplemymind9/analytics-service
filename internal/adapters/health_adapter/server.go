package health_adapter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/cripplemymind9/analytics-service/pkg/pb/health"
)

type Service struct {
	pb.UnimplementedHealthServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterHealthServiceServer(server, s)
}

func (s *Service) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterHealthServiceHandler(ctx, mux, conn)
}