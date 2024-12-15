package adapter

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ImplementationAdapter interface {
	RegisterServer(grpcServer *grpc.Server)
	RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}