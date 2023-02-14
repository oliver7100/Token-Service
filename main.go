package main

import (
	"net"

	"github.com/oliver7100/token-service/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterAuthServiceServer(
		s,
		proto.NewService(),
	)

	panic(s.Serve(listener))
}
