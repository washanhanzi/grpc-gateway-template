package controller

import (
	"context"

	"github.com/infiniteprimes/grpc-gateway-template/example"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) Echo(ctx context.Context, in *example.StringMessage) (*example.StringMessage, error) {
	return &example.StringMessage{Value: in.Value}, nil
}
