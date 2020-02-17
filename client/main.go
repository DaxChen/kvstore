package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"runtime"
	"sync"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosuri/uiprogress"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	command    = flag.String("command", "", "load/exp1/exp2/stat/getPrefix")
	mode       = flag.String("mode", "", "read/readwrite")
	numKeys    = flag.Int("num_keys", 0, "number of keys")
	valueSize  = flag.Int("value_size", 0, "value size in bytes")
	prefixKey  = flag.String("prefix_key", "", "the prefix key to get")
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomValue(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func doGet(client pb.KVStoreClient, key string) bool {
	log.Tracef("try calling Get(%s)", key)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	value, err := client.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		log.Debugf("called Get, got error %v", err)
		return false
	}
	log.Tracef("called Get(%s), got %s", key, value)
	return true
}

func doSet(client pb.KVStoreClient, key string, value string) bool {
	log.Tracef("try calling Set(%s, %s)", key, value)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.Set(ctx, &pb.KeyValuePair{Key: key, Value: value})
	if err != nil {
		log.Errorf("called Set, got error %v", err)
		return false
	}
	log.Tracef("called Set(%s, %s), got %v", key, value, res)
	return true
}

func doSetWithoutSync(client pb.KVStoreClient, key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.SetWithoutSync(ctx, &pb.KeyValuePair{Key: key, Value: value})
	if err != nil {
		log.Errorf("called SetWithoutSync, got error %v", err)
	}
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

func doStat(client pb.KVStoreClient) {
	log.Tracef("try calling GetStat")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	states, err := client.GetStat(ctx, &empty.Empty{})
	if err != nil {
		log.Errorf("called GetStat, got error %v", err)
	}

	log.Debug("time", states.ServerStartTime, "\t#gets", states.TotalGetsDone,
		"\t#sets", states.TotalSetsDone, "\t#prefixs", states.TotalGetprefixesDone)
}

func doCrash(client pb.KVStoreClient) {
	log.Debug("try calling Crash()")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client.Crash(ctx, &empty.Empty{})
}

func loadDataBase(client pb.KVStoreClient, numKeys int, valueSize int) {
	log.Infof("load Data: %d keys, each %d bytes value.", numKeys, valueSize)

	// progress bar
	bar := uiprogress.AddBar(numKeys).AppendCompleted().PrependElapsed()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%d/%d", b.Current(), numKeys)
	})
	uiprogress.Start()

	// fan-in
	in := make(chan int)
	go func() {
		for i := 1; i <= numKeys; i++ {
			in <- i
		}
		close(in)
	}()

	var wg sync.WaitGroup
	for worker := 0; worker < runtime.NumCPU(); worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range in {
				key := fmt.Sprintf("%0128d", i)
				value := generateRandomValue(valueSize)

				doSetWithoutSync(client, key, value)
				bar.Incr()
			}
		}()
	}
	wg.Wait()
	uiprogress.Stop()

	// call fsync
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Fsync(ctx, &empty.Empty{})
	if err != nil {
		log.Errorf("called FSync, got error %v", err)
	}
}

func getAveReadLatency(client pb.KVStoreClient, numKeys int) {
	// throughput (operations per second) = (total number of operations) / time.
	// latency of the run as the sum of latencies of all operations divided by the total no. of operations.
	log.Debug("getAveReadLatency")
	var dur time.Duration = 0
	countGet := 0

	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if countGet == 0 {
					log.Debug("no get")
				} else {
					log.Debug("latency ", dur/time.Duration(countGet), ", throughput(ops/sec) ", float64(countGet)/dur.Seconds())
				}
			}
		}
	}()

	exp := time.Now()
	for time.Now().Before(exp.Add(3 * time.Minute)) {
		key := fmt.Sprintf("%0128d", rand.Intn(numKeys)+1)
		period, suc := getReadLatency(client, key)
		if suc {
			countGet++
		}
		dur += period
	}
	ticker.Stop()
	done <- true

	if countGet == 0 {
		log.Info("no get")
	} else {
		log.Info("latency ", dur/time.Duration(countGet), ", throughput(ops/sec) ", float64(countGet)/dur.Seconds())
	}
}

func getAveRWLatency(client pb.KVStoreClient, numKeys int, valueSize int) {
	log.Debug("getAveRWLatency")
	var durR time.Duration = 0
	var durW time.Duration = 0
	countGet := 0
	countSet := 0

	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if countGet == 0 && countSet == 0 {
					log.Debug("no get and set")
				} else {
					log.Debug("readLatency ", durR/time.Duration(countGet),
						", writeLatency ", durW/time.Duration(countSet),
						", throughput(ops/sec) ", float64(countGet+countSet)/(durR.Seconds()+durW.Seconds()))
				}
			}
		}
	}()

	exp := time.Now()
	for time.Now().Before(exp.Add(3 * time.Minute)) {
		key := fmt.Sprintf("%0128d", rand.Intn(numKeys)+1)
		if rand.Intn(2) >= 1 {
			period, suc := getReadLatency(client, key)
			if suc {
				countGet++
			}
			durR += period
		} else {
			value := generateRandomValue(valueSize)
			period, suc := getWriteLatency(client, key, value)
			if suc {
				countSet++
			}
			durW += period
		}
	}
	ticker.Stop()
	done <- true

	if countGet == 0 && countSet == 0 {
		log.Info("no get and set")
	} else {
		log.Info("readLatency ", durR/time.Duration(countGet),
			", writeLatency ", durW/time.Duration(countSet),
			", throughput(ops/sec) ", float64(countGet+countSet)/(durR.Seconds()+durW.Seconds()))
	}
}

func getReadLatency(client pb.KVStoreClient, key string) (time.Duration, bool) {
	start := time.Now()
	suc := doGet(client, key)
	end := time.Now()

	return end.Sub(start), suc
}

func getWriteLatency(client pb.KVStoreClient, key string, value string) (time.Duration, bool) {
	start := time.Now()
	suc := doSet(client, key, value)
	end := time.Now()

	return end.Sub(start), suc
}

func getColdLatency(client pb.KVStoreClient) {
	log.Debug("getColdLatency")
	go doCrash(client)

	exp2 := time.Now()
	var start time.Time
	fail := 0
	for time.Now().Before(exp2.Add(time.Minute)) {
		key := fmt.Sprintf("%0128d", 1)
		suc := doGet(client, key)
		if !suc && fail == 0 {
			log.Debug("fail to connect", time.Now())
			fail++
			start = time.Now()
		}
		if suc && fail != 0 {
			log.Debug("recover connection", time.Now())
			break
		}
	}

	log.Debug("cold latency: ", time.Now().Sub(start))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use all available cpu cores
	log.SetLevel(log.InfoLevel)

	flag.Parse()

	log.Infof("trying to connect to server at %s", *serverAddr)
	const maxMsgSize = 1024 * 1024 * 5
	conn, err := grpc.Dial(
		*serverAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(maxMsgSize),
			grpc.MaxCallSendMsgSize(maxMsgSize),
		),
	)
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	log.Infof("connected to %s", *serverAddr)
	defer conn.Close()

	client := pb.NewKVStoreClient(conn)

	// <load> <#keys> <valueSize>
	// <exp1> <read> <#Keys>
	// <exp1> <readwrite> <#Keys> <valueSize>
	// <exp2>
	// <stat>
	// <prefix> <prefixKey>
	if *command == "load" {
		loadDataBase(client, *numKeys, *valueSize)
	} else if *command == "exp1" {
		if *mode == "read" {
			getAveReadLatency(client, *numKeys)
		} else if *mode == "readwrite" {
			getAveRWLatency(client, *numKeys, *valueSize)
		}
	} else if *command == "exp2" {
		getColdLatency(client)
	} else if *command == "stat" {
		doStat(client)
	} else if *command == "prefix" {
		doGetPrefix(client, *prefixKey)
	} else {
		flag.PrintDefaults()
		fmt.Println("")
		color.Red("Usage:")
		fmt.Println("# load the server with <#keys> keys, each with <value_size> bytes of value.")
		color.Green("./kvclient -server_addr=localhost:10000 -command=load -num_keys=<#keys> -value_size=<value_size>")
		fmt.Println("")
		fmt.Println("# run exp1 read test (provide #keys to make sure randome genreate keys are in range)")
		color.Green("./kvclient -command=exp1 -mode=read -num_keys=<#keys>")
		fmt.Println("# run exp1 50% read 50% write test")
		color.Green("./kvclient -command=exp1 -mode=readwrite -num_keys=<#keys> -value_size=<value_size>")
		fmt.Println("")
		fmt.Println("# run exp2: measure cold start time")
		color.Green("./kvclient -command=exp2")
		fmt.Println("")
		fmt.Println("# other commands")
		color.Green("./kvclient -command=stat")
		color.Green("./kvclient -command=prefix -prefix_key=<prefix_key>")
	}
}
