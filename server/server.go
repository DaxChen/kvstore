package main

import (
	"context"
	"os"

	pb "github.com/DaxChen/kvstore/proto"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
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
	log.Tracef("received request Get(%s)\n", k)

	value, err := s.store.Get(k)
	if err != nil {
		return nil, err
	}
	return &pb.Value{Value: value}, nil
}

func (s *Server) Set(ctx context.Context, pair *pb.KeyValuePair) (*pb.SetResponse, error) {
	k, v := pair.GetKey(), pair.GetValue()
	log.Tracef("received request Set, total set: %d", totalSetsDone)

	s.store.Set(k, v, true)

	return &pb.SetResponse{Success: true}, nil
}

func (s *Server) SetWithoutSync(ctx context.Context, pair *pb.KeyValuePair) (*pb.SetResponse, error) {
	k, v := pair.GetKey(), pair.GetValue()
	log.Tracef("received request SetWithoutSync, total set: %d", totalSetsDone)

	s.store.Set(k, v, false)

	return &pb.SetResponse{Success: true}, nil
}

func (s *Server) Fsync(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if err := s.store.logfile.Sync(); err != nil {
		log.Errorf("Fsync: error fsync to logfile: %v", err)
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *Server) GetPrefix(prefixKey *pb.PrefixKey, stream pb.KVStore_GetPrefixServer) error {
	prefix := prefixKey.GetPrefix()
	log.Debugf("received request GetPrefix(%s)\n", prefix)

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

func (s *Server) GetStat(context.Context, *empty.Empty) (*pb.States, error) {
	startTime, getsDone, setsDone, prefixesDone := s.store.GetStats()
	//log.Debugf("server getsDone(%d)\n", getsDone)
	return &pb.States{ServerStartTime: startTime, TotalGetsDone: getsDone,
		TotalSetsDone: setsDone, TotalGetprefixesDone: prefixesDone}, nil
}

func (s *Server) Crash(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	os.Exit(3)
	return &empty.Empty{}, nil
}
