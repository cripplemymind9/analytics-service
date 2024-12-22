package auth_service

import (
	"context"

	"github.com/cripplemymind9/analytics-service/pkg/pb/auth"
)

type Auth interface {
	GenerateToken(ctx context.Context, req *auth.LoginRequest) (string, error)
	CreateUser(ctx context.Context, req *auth.SigninRequest) (string, error)
}