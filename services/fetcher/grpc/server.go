package grpc

import (
	"context"

	pb "github.com/hatena/Hatena-Intern-2021/services/fetcher/pb/fetcher"
	fetcher "github.com/hatena/Hatena-Intern-2021/services/fetcher/src"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.FetcherServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetch は受け取った Link を Title に変換する
func (s *Server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchReply, error) {
	title, err := fetcher.Fetch(ctx, in.Link)
	if err != nil {
		return nil, err
	}
	return &pb.FetchReply{Title: title}, nil
}
