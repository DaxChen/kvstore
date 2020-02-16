package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
)

type Store struct {
	cache   *sync.Map
	logfile *os.File
}

// NewStore Constructor for Store
func NewStore(filename string) *Store {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file %s, error: %v", filename, err)
	}

	var cache sync.Map
	// TODO: populate cache with file upon recover
	readFile, err := os.Open(filename)
	if err != nil {
		log.Warnf("Error opening logfile to read, perhaps logfile not exist yet? %v", err)
	} else {
		defer readFile.Close()

		reader := bufio.NewReaderSize(readFile, 512*1024)
		for {
			var (
				isPrefix   bool  = true
				err        error = nil
				line, part []byte
			)
			for isPrefix && err == nil {
				part, isPrefix, err = reader.ReadLine()
				line = append(line, part...)
				// log.Tracef("readine parts for line %d, part len %d, line len %d\n", linenum, len(part), len(line))
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Errorf("Error from reader while reading logfile: %v", err)
			}

			//log.Debugf("reading logfile: %s", line)
			var data map[string]interface{}
			err = json.Unmarshal(line, &data)
			if err != nil {
				log.Errorf("Error from json.Unmarshal: %v", err)
			}
			key := data["key"].(string)
			value := data["value"].(string)
			//log.Debugf("Unmarshal json: %s, %s", key, value)
			cache.Store(key, value)
		}
	}

	return &Store{
		cache:   &cache,
		logfile: f,
	}
}

func (s *Store) Get(key string) (string, error) {
	if result, ok := s.cache.Load(key); ok {
		return result.(string), nil
	}
	return "", errors.New("key not found")
}

func (s *Store) Set(key, value string) error {
	// first append to logfile
	payload, err := json.Marshal(map[string]string{"key": key, "value": value})
	if err != nil {
		log.Fatalf("error in json.Marshal(key, value): %v", err)
	}
	if _, err := s.logfile.Write(append(payload, '\n')); err != nil {
		s.logfile.Close() // ignore error; Write error takes precedence
		log.Fatalf("error appending to logfile: %v", err)
		return err
	}
	// call fsync to make sure commit to file
	if err := s.logfile.Sync(); err != nil {
		log.Fatalf("error fsync to logfile: %v", err)
		return err
	}

	// finally add to cache
	s.cache.Store(key, value)
	return nil
}

func (s *Store) GetPrefix(prefix string, callback func(key, value string) bool) {
	s.cache.Range(func(k, v interface{}) bool {
		if !strings.HasPrefix(k.(string), prefix) {
			return true
		}

		return callback(k.(string), v.(string))
	})
}
