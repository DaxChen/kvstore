package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	pb "github.com/DaxChen/kvstore/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosuri/uiprogress"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	value, err := client.Get(ctx, &pb.Key{Key: key})
	if err != nil {
		//log.Errorf("called Get(%s), got error %v", key, err)
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
}

func getAveReadLatency(client pb.KVStoreClient, numKeys int) {
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
					log.Debug("latency ", dur/time.Duration(countGet), " ,throughput ", float64(128*countGet/1024)/dur.Seconds())
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
		log.Info("latency ", dur/time.Duration(countGet), " ,throughput ", float64(128*countGet/1024)/dur.Seconds())
	}
}

func getAveRWLatency(client pb.KVStoreClient, numKeys int, valueSize int) {
	log.Debug("getAveRWLatency")
	var dur time.Duration = 0
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
					rThp := float64(128*countGet/1024) / dur.Seconds()
					wThp := float64(valueSize*countSet/1024) / dur.Seconds()
					log.Debug("latency ", dur/time.Duration(countGet+countSet), " ,throughput ", rThp+wThp)
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
			dur += period
		} else {
			value := generateRandomValue(valueSize)
			period, suc := getWriteLatency(client, key, value)
			if suc {
				countSet++
			}
			dur += period
		}
	}
	ticker.Stop()
	done <- true

	if countGet == 0 && countSet == 0 {
		log.Info("no get and set")
	} else {
		rThp := float64(128*countGet/1024) / dur.Seconds()
		wThp := float64(valueSize*countSet/1024) / dur.Seconds()
		log.Info("latency ", dur/time.Duration(countGet+countSet), " ,throughput ", rThp+wThp)
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
	log.SetLevel(log.DebugLevel)

	log.Info("trying to connect to server at localhost:10000")
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	log.Info("connected to localhost:10000")
	defer conn.Close()

	client := pb.NewKVStoreClient(conn)

	args := os.Args[1:]
	var numKeys int64
	var valueSize int64
	// <load> <#keys> <valueSize>
	// <exp1> <read> <#Keys>
	// <exp1> <readwrite> <#Keys> <valueSize>
	// <exp2>
	// <stat>
	// <prefix> <prefixKey>
	if len(args) == 3 && args[0] == "load" {
		numKeys, err = strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Debugf("fail to parse numKeys")
		}
		valueSize, err = strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Debugf("fail to parse valueSize")
		}

		loadDataBase(client, int(numKeys), int(valueSize))
	} else if len(args) == 3 && args[0] == "exp1" {
		if args[1] == "read" {
			numKeys, err = strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				log.Debugf("fail to parse numKeys")
			}

			getAveReadLatency(client, int(numKeys))
		}
	} else if len(args) == 4 && args[0] == "exp1" {
		if args[1] == "readwrite" {
			numKeys, err = strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				log.Debugf("fail to parse numKeys")
			}
			valueSize, err = strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				log.Debugf("fail to parse valueSize")
			}

			getAveRWLatency(client, int(numKeys), int(valueSize))
		}
	} else if len(args) == 1 && args[0] == "exp2" {
		getColdLatency(client)
	} else if len(args) == 1 && args[0] == "stat" {
		doStat(client)
	} else if len(args) == 2 && args[0] == "prefix" {
		doGetPrefix(client, args[1])
	}
	//else {
	//	fmt.Println("")
	//	color.Red("Usage:")
	//	fmt.Println("# load the server with <#keys> keys, each with <value_size> bytes of value.")
	//	color.Green("./kvclient load <#keys> <value_size>")
	//	fmt.Println("")
	//	fmt.Println("# run exp1 read test (provide #keys to make sure randome genreate keys are in range)")
	//	color.Green("./kvclient exp1 read <#keys>")
	//	fmt.Println("# run exp1 50% read 50% write test")
	//	color.Green("./kvclient exp1 readwrite <#keys> <value_size>")
	//	fmt.Println("")
	//	fmt.Println("# run exp2: measure cold start time")
	//	color.Green("./kvclient exp2")
	//	fmt.Println("")
	//	fmt.Println("# other commands")
	//	color.Green("./kvclient stat")
	//	color.Green("./kvclient prefix <prefix_key>")
	//}
}
