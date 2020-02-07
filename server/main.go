package main

import (
	"context"
	"errors"
	"net"
	"strings"
	"sync"

	pb "github.com/DaxChen/kvstore/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct{}

// var cache map[string]string = make(map[string]string)
var cache sync.Map

func main() {
	log.SetLevel(log.TraceLevel)

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Info("starting server on port 10000...")
	grpcServer := grpc.NewServer()
	pb.RegisterKVStoreServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

func (s *server) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	k := key.GetKey()
	log.Debugf("received request Get(%s)\n", k)

	// v = cache[k]
	if result, ok := cache.Load(k); ok {
		v := result.(string)
		return &pb.Value{Value: v}, nil
	}
	return nil, errors.New("key not found")
}

func (s *server) Set(ctx context.Context, pair *pb.KeyValuePair) (*pb.SetResponse, error) {
	k, v := pair.GetKey(), pair.GetValue()
	log.Debugf("received request Set(%s, %s)\n", k, v)

	// cache[k] = v
	cache.Store(k, v)

	return &pb.SetResponse{Success: true}, nil
}

func (s *server) GetPrefix(prefixKey *pb.PrefixKey, stream pb.KVStore_GetPrefixServer) error {
	prefix := prefixKey.GetPrefix()
	log.Debugf("received request GetPrefix(%s)\n", prefix)

	var streamError error
	cache.Range(func(k, v interface{}) bool {
		if !strings.HasPrefix(k.(string), prefix) {
			return true
		}
		if err := stream.Send(&pb.Value{Value: v.(string)}); err != nil {
			streamError = err
			return false
		}
		return true
	})

	if streamError != nil {
		return streamError
	}
	return nil
}
