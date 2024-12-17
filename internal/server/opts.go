package server

import (
	"google.golang.org/grpc"
	"github.com/cripplemymind9/analytics-service/internal/server/adapter"
)

// Структура содержит натстройки для сервера, такие как адреса и адаптеры 
type options struct {
	adapters 						[]adapter.ImplementationAdapter
	grpcUnaryServerInterceptors		[]grpc.UnaryServerInterceptor
	grpcAddress 					string
	httpAddress						string
}

type option func(o *options)

func (o option) apply(os *options) { o(os) }

func WithImplementationAdapters(adapters ...adapter.ImplementationAdapter) EntrypointOption {
	return option(func(o *options) { o.adapters = append(o.adapters, adapters...) })
}

func WithGrpcUnaryServerInterceptors(
	grpcUnaryServerInterceptors ...grpc.UnaryServerInterceptor,
) EntrypointOption {
	return option(func(o *options) {
		o.grpcUnaryServerInterceptors = append(o.grpcUnaryServerInterceptors, grpcUnaryServerInterceptors...)
	})
}