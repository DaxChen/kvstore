package main

import (
	"net"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const maxMsgSize = 1024 * 1024 * 5

func main() {
	log.SetLevel(log.TraceLevel)

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Info("starting server on port 10000...")
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(maxMsgSize), grpc.MaxSendMsgSize(maxMsgSize))
	server := NewServer()
	pb.RegisterKVStoreServer(grpcServer, server)
	log.Info("server started on time: ", time.Now())
	grpcServer.Serve(lis)
}
