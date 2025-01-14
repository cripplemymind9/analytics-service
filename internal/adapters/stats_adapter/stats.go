package stats_adapter

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/cripplemymind9/analytics-service/pkg/pb/stats"
)

func (s *Service) GetStats(ctx context.Context, req *stats.GetStatsRequest) (*stats.GetStatsResponse, error) {
	if req.GetFrom() == "" || req.GetTo() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "'from' and 'to' parameters cannot be empty")
	}

	response, err := s.svc.GetStats(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get stats: %v", err)
	}

	return response, nil
}
