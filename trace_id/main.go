package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTraceID(length int64) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result string
	for i := 0; i < int(length); i++ {
		result += string(charset[rand.Intn(len(charset))])
	}

	return result
}

func main() {
	for i :=0; i < 10000; i++ {
		fmt.Println(GenerateTraceID(6))
	}

}