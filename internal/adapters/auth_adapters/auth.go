package auth_adapter

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/cripplemymind9/analytics-service/pkg/pb/auth"
)

func (s *Service) Login(
	ctx context.Context,
	req *auth.LoginRequest,
) (*auth.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "All fields (email, password) are required")
	}

	token, err := s.svc.GenerateToken(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate token: %v", err)
	}

	return &auth.LoginResponse{
		Token: token,
	}, nil
}

func (s *Service) Signin(
	ctx context.Context,
	req *auth.SigninRequest,
) (*auth.SigninResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "All fields (email, password) are required")
	}

	token, err := s.svc.CreateUser(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed create user: %v", err)
	}

	return &auth.SigninResponse{
		Token: token,
	}, nil
}