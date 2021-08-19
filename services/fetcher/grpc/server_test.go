package grpc

import (
	"context"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2021/services/fetcher/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetch(t *testing.T) {
	s := NewServer()
	link := "https://hatenacorp.jp/recruit/intern/2021"
	reply, err := s.Fetch(context.Background(), &pb.FetchRequest{Link: link})
	assert.NoError(t, err)
	assert.Equal(t, "はてなリモートインターンシップ2021 - 株式会社はてな", reply.Title)
}
