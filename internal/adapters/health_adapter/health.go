package health_adapter

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CheckHealth(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
