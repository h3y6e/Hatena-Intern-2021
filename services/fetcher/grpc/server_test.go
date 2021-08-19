package grpc

import (
	"context"
	"testing"

	pb "github.com/hatena/Hatena-Intern-2021/services/fetcher/pb/fetcher"
	"github.com/stretchr/testify/assert"
)

func Test_Server_Fetch(t *testing.T) {
	s := NewServer()
	link := `https://google.com/`
	reply, err := s.Fetch(context.Background(), &pb.FetchRequest{Link: link})
	assert.NoError(t, err)
	assert.Equal(t, "title", reply.Title)
}
