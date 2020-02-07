package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DaxChen/kvstore/proto"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("starting server on port 10000...")
	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

func (s *server) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	k := key.GetKey()
	log.Printf("received request Get(%s)\n", k)

	return &pb.Value{Value: "Value of " + k}, nil
}

func (s *server) Set(ctx context.Context, pair *pb.KeyValuePair) (*pb.SetResponse, error) {
	k, v := pair.GetKey(), pair.GetValue()
	log.Printf("received request Set(%s, %s)\n", k, v)
	return &pb.SetResponse{Success: true}, nil
}

func (s *server) GetPrefix(prefix *pb.PrefixKey, stream pb.KVStore_GetPrefixServer) error {
	log.Printf("received request GetPrefix(%s)\n", prefix.GetPrefix())
	return nil
}
