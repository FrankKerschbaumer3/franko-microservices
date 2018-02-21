package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type issueFetcher struct{}

func (*issueFetcher) FetchIssues(context.Context, *FetchReq) (*FetchReply, error) {
	return &FetchReply{Issues: nil}, nil
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
	RegisterIssueFetcherServer(grpcServer, &issueFetcher{})
	grpcServer.Serve(lis)
}
