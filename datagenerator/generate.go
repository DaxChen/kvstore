package main

import (
	"fmt"
	"math/rand"
	"os"
)

var (
	f512 *os.File
	f4KB *os.File
	f512KB *os.File
	f1MB *os.File
	f4MB *os.File
	keys *os.File

	mix512 *os.File
	mix4KB *os.File
	mix512KB *os.File
	mix1MB *os.File
	mix4MB *os.File

	err error
)

func main() {
	keys, err = os.Create("./keys.txt")
	f512, err = os.Create("./512v1.txt")
	f4KB, err = os.Create("./4KBv1.txt")
	f512KB, err = os.Create("./512KBv1.txt")
	f1MB, err = os.Create("./1MBv1.txt")
	f4MB, err = os.Create("./4MBv1.txt")

	mix512, err = os.Create("./mix512.txt")
	mix4KB, err = os.Create("./mix4KB.txt")
	mix512KB, err = os.Create("./mix512KB.txt")
	mix1MB, err = os.Create("./mix1MB.txt")
	mix4MB, err = os.Create("mix4MB.txt")

	var i int
	for i = 0; i < 1000; i++ {
		if i % 100 == 0 {
			fmt.Printf("finished %d", i)
		}

		key := StringKey(128)
		v512 := StringValue(512)
		v4KB := StringValue(4000)
		v512KB := StringValue(512000)
		v1MB := StringValue(1000000)
		v4MB := StringValue(4000000)

		keys.WriteString(key + "\n")
		f512.WriteString(key + " " + v512 + "\n")
		f4KB.WriteString(key + " " + v4KB + "\n")
		f512KB.WriteString(key + " " + v512KB + "\n")
		f1MB.WriteString(key + " " + v1MB + "\n")
		f4MB.WriteString(key + " " + v4MB + "\n")

		if rand.Intn(2) >= 1 {
			mix512.WriteString(key + "\n")
			mix4KB.WriteString(key + "\n")
			mix512KB.WriteString(key + "\n")
			mix1MB.WriteString(key + "\n")
			mix4MB.WriteString(key + "\n")
		} else {
			mix512.WriteString(key + " " + v512 + "\n")
			mix4KB.WriteString(key + " " + v4KB + "\n")
			mix512KB.WriteString(key + " " + v512KB + "\n")
			mix1MB.WriteString(key + " " + v1MB + "\n")
			mix4MB.WriteString(key + " " + v4MB + "\n")
		}
	}

	keys.Sync()
	f512.Sync()
	f4KB.Sync()
	f512KB.Sync()
	f1MB.Sync()
	f4MB.Sync()

	mix512.Sync()
	mix4KB.Sync()
	mix512KB.Sync()
	mix1MB.Sync()
	mix4MB.Sync()

	keys.Close()
	f512.Close()
	f4KB.Close()
	f512KB.Close()
	f1MB.Close()
	f4MB.Close()

	mix512.Close()
	mix4KB.Close()
	mix512KB.Close()
	mix1MB.Close()
	mix4MB.Close()
}
