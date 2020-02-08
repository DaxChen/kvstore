package main

import (
	"net"

	pb "github.com/DaxChen/kvstore/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.SetLevel(log.TraceLevel)

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Info("starting server on port 10000...")
	grpcServer := grpc.NewServer()
	server := NewServer()
	pb.RegisterKVStoreServer(grpcServer, server)
	log.Info("server started")
	grpcServer.Serve(lis)
}
