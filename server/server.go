package main

import (
	"context"

	pb "github.com/DaxChen/kvstore/proto"
)

type Server struct {
	store *Store
}

// NewServer Constructor of Server
func NewServer() *Server {
	store := NewStore("mem.log")
	return &Server{
		store: store,
	}
}

func (s *Server) Get(ctx context.Context, key *pb.Key) (*pb.Value, error) {
	k := key.GetKey()
	//log.Debugf("received request Get(%s)\n", k)

	value, err := s.store.Get(k)
	if err != nil {
		return nil, err
	}
	return &pb.Value{Value: value}, nil
}

func (s *Server) Set(ctx context.Context, pair *pb.KeyValuePair) (*pb.SetResponse, error) {
	k, v := pair.GetKey(), pair.GetValue()
	//log.Debugf("received request Set(%s, %s)\n", k, v)

	// cache[k] = v
	s.store.Set(k, v)

	return &pb.SetResponse{Success: true}, nil
}

func (s *Server) GetPrefix(prefixKey *pb.PrefixKey, stream pb.KVStore_GetPrefixServer) error {
	prefix := prefixKey.GetPrefix()
	//log.Debugf("received request GetPrefix(%s)\n", prefix)

	var streamError error

	s.store.GetPrefix(prefix, func(key, value string) bool {
		if err := stream.Send(&pb.Value{Value: value}); err != nil {
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
