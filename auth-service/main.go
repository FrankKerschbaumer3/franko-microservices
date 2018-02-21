package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Authentication Service impl
type authService struct{}

func (*authService) Authenticate(ctx context.Context, req *AuthRequest) (*AuthReply, error) {
	return &AuthReply{IsAuthenticated: true}, nil
}

func main() {
	port := flag.Int("port", 13500, "RPC server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterAuthServer(grpcServer, &authService{})
	grpcServer.Serve(lis)
}
