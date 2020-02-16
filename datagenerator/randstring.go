package main

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func StringValue(length int) string {
	return StringWithCharset(length, charset)
}

func StringKey(length int) string {
	b := make([]byte, length)
	index := rand.Intn(len(charset))
	for i := range b {
		b[i] = charset[index]
	}
	return string(b)
}