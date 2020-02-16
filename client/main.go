package main

import (
	"bufio"
	"context"
	"io"
	"os"
	"strings"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const loop = 1

func doGet(client pb.KVStoreClient, key string) bool {
	log.Tracef("try calling Get(%s)", key)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		//log.Errorf("called Get(%s), got error %v", key, err)
		return false
	}
	//log.Debugf("called Get(%s), got %s", key, value)
	return true
}

func doSet(client pb.KVStoreClient, key string, value string) bool {
	log.Tracef("try calling Set(%s, %s)", key, value)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Set(ctx, &pb.KeyValuePair{Key: key, Value: value})
	if err != nil {
		//log.Errorf("called Set(%s, %s), got error %v", key, value, err)
		return false
	}
	//log.Debugf("called Set(%s, %s), got %v", key, value, res)
	return true
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
	count := 0

	for {
		value, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("called GetPrefix(%s), got error %v", prefix, err)
		}
		count++
		values = append(values, value.GetValue())
	}
	log.Debugf("called GetPrefix(%s), got %v, total get prefixes done %d", prefix, values, count)
}

func getAveLatency(client pb.KVStoreClient, filename string) {
	for i := 0; i < loop; i++ {
		log.Tracef("load file %s", filename)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("failed to open file: %s", filename)
		}
		defer file.Close()

		var dur time.Duration
		countGet := 0
		countSet := 0

		scanner := bufio.NewScanner(file)
		scanner.Buffer([]byte{}, 1024*1024*5)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			kv := strings.Split(scanner.Text(), " ")
			if len(kv) == 1 {
				period, suc := getReadLatency(client, kv[0])
				if suc {
					countGet++
				}
				dur += period
			} else {
				period, suc := getWirteLatency(client, kv[0], kv[1])
				if suc {
					countSet++
				}
				dur += period
			}
		}

		log.Println("time: ", time.Now(), " ,latency: ", dur / time.Duration(countGet + countSet),
			" ,total gets done: ", countGet, " ,total sets done: ", countSet)
	}
}

func getReadLatency(client pb.KVStoreClient, key string) (time.Duration, bool) {
	start := time.Now()
	suc := doGet(client, key)
	end := time.Now()

	return end.Sub(start), suc
}

func getWirteLatency(client pb.KVStoreClient, key string, value string) (time.Duration, bool) {
	start := time.Now()
	suc := doSet(client, key, value)
	end := time.Now()

	return end.Sub(start), suc
}

func main() {
	log.SetLevel(log.DebugLevel)

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	log.Info("connected to localhost:10000")
	defer conn.Close()

	client := pb.NewKVStoreClient(conn)


	args := os.Args[1:]
	// "./client/main.exe" "exp1" "f512.txt"
	// "./client/main.exe" "exp2" "key"
	// "./client/main.exe" "prefix" "A"
	if len(args) == 2 && strings.Compare(args[0], "exp1") == 0 {
		getAveLatency(client, args[1])
	} else if len(args) == 2 && strings.Compare(args[0], "exp2") == 0 {
		exp2 := time.Now()

		var start time.Time
		fail := 0
		for exp2.Before(exp2.Add(time.Minute)) {
			suc := doGet(client, args[1])

			if !suc && fail == 0 {
				log.Info("fail on", time.Now())
				fail++
				start = time.Now()
			}
			if suc && fail != 0 {
				log.Info("recover on", time.Now())
				break
			}
		}

		log.Info("cold latency: ", time.Now().Sub(start))
	} else if len(args) == 2 && strings.Compare(args[0], "prefix") == 0 {
		doGetPrefix(client, args[1])
	}

}
