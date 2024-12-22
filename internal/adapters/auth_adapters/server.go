package auth_adapter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/auth_service"
	pb "github.com/cripplemymind9/analytics-service/pkg/pb/auth"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	svc auth_service.Auth
}

func New(svc auth_service.Auth) *Service {
	return &Service{svc: svc}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterAuthServiceServer(server, s)
}

func (s *Service) RegisterHandler(
	ctx context.Context,
	mux *runtime.ServeMux,
	conn *grpc.ClientConn,
) error {
	return pb.RegisterAuthServiceHandler(ctx, mux, conn)
}
