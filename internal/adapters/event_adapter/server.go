package event_adapter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/analytics-service/internal/interfaces/service/event_service"
	pb "github.com/cripplemymind9/analytics-service/pkg/pb/event"
)

type Service struct {
	pb.UnimplementedEventServiceServer
	svc event_service.Event
}

func New(svc event_service.Event) *Service {
	return &Service{svc: svc}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterEventServiceServer(server, s)
}

func (s *Service) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterEventServiceHandler(ctx, mux, conn)
}
