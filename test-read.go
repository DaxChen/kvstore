package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	readFile, err := os.Open("test.log")
	if err != nil {
		log.Fatalf("Error opening logfile to read, perhaps logfile not exist yet? %v", err)
	} else {
		defer readFile.Close()

		reader := bufio.NewReaderSize(readFile, 128*1024)
		linenum := 1
		for {
			var (
				isPrefix bool   = true
				err      error  = nil
				line     string = ""
				part     []byte
			)
			for isPrefix && err == nil {
				part, isPrefix, err = reader.ReadLine()
				line += string(part)
				log.Printf("readine parts for line %d, part len %d, line len %d\n", linenum, len(part), len(line))
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error from reader while reading logfile: %v\n", err)
			}

			log.Printf("reading logfile, got %d chars for line %d\n", len(line), linenum)

			linenum++
		}
	}
}
