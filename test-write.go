package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file %s, error: %v", "test.log", err)
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 512*1024; j++ {
			if _, err := file.Write([]byte("g")); err != nil {
				file.Close() // ignore error; Write error takes precedence
				log.Fatalf("error appending to logfile: %v", err)
				return
			}
		}
		if _, err := file.Write([]byte("\n")); err != nil {
			file.Close() // ignore error; Write error takes precedence
			log.Fatalf("error appending to logfile: %v", err)
			return
		}
		// call fsync to make sure commit to file
		if err := file.Sync(); err != nil {
			log.Fatalf("error fsync to logfile: %v", err)
			return
		}
	}
}
