package main

import (
	"context"
	"io"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func doGet(client pb.KVStoreClient, key string) {
	log.Tracef("try calling Get(%s)", key)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	value, err := client.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		log.Errorf("called Get(%s), got error %v", key, err)
	}
	log.Debugf("called Get(%s), got %s", key, value)
}

func doSet(client pb.KVStoreClient, key string, value string) {
	log.Tracef("try calling Set(%s, %s)", key, value)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Set(ctx, &pb.KeyValuePair{Key: key, Value: value})
	if err != nil {
		log.Errorf("called Set(%s, %s), got error %v", key, value, err)
	}
	log.Debugf("called Set(%s, %s), got %v", key, value, res)
}

func doGetPrefix(client pb.KVStoreClient, prefix string) {
	log.Tracef("try calling GetPrefix(%s)", prefix)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GetPrefix(ctx, &pb.PrefixKey{Prefix: prefix})
	if err != nil {
		log.Errorf("called GetPrefix(%s), got error %v", prefix, err)
	}

	var values []string
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("called GetPrefix(%s), got error %v", prefix, err)
		}
		values = append(values, value.GetValue())
	}
	log.Debugf("called GetPrefix(%s), got %v", prefix, values)
}

func main() {
	log.SetLevel(log.DebugLevel)

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	log.Info("Connected to localhost:10000")
	defer conn.Close()

	client := pb.NewKVStoreClient(conn)

	doSet(client, "abc", "123")
	doSet(client, "def", "456")
	doSet(client, "ghi", "789")
	doSet(client, "jkl", "012")
	doGet(client, "abc")
	doGet(client, "def")
	doGet(client, "ghi")
	doGet(client, "jkl")
	doSet(client, "abc", "abc")
	doSet(client, "apple", "123123apple")
	doSet(client, "app", "0000app")
	doGetPrefix(client, "app")
	doGetPrefix(client, "a")
}
