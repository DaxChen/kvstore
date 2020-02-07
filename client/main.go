package main

import (
	"context"
	"log"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	"google.golang.org/grpc"
)

func doGet(client pb.KVStoreClient, key string) {
	log.Println("try calling Get with", key)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	value, err := client.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		log.Fatalf("%v.Get(_) = _, %v: ", client, err)
	}
	log.Println(value)
}

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewKVStoreClient(conn)

	// try calling Get
	doGet(client, "abc")
}
