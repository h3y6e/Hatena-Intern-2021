package grpc

import (
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.AccountServer に対する実装
type Server struct {
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}
