package main

import (
	"flag"
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/MatthewEdge/franko-services/auth-service/auth"
)

// Authentication Service impl
type authService struct {}

func (*authService) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	return &pb.AuthReply{}, nil
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
	pb.RegisterAuthServer(grpcServer, &authService{})
	grpcServer.Serve(lis)
}