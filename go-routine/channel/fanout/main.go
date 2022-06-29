package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var publisherID int64
var workerID int64

func main() {
	input := make(chan string)
	go worker(input)
	go worker(input)
	go publisher(input)
	go publisher(input)
	time.Sleep(1 * time.Second)
}

func worker(in <-chan string) {
	atomic.AddInt64(&workerID, 1)
	thisID := atomic.LoadInt64(&workerID)
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		fmt.Printf("%d: input is: %s\n", thisID, <-in)
	}
}

func publisher(out chan<- string) {
	atomic.AddInt64(&publisherID, 1)
	thisID := atomic.LoadInt64(&publisherID)
	dataID := 0
	for {
		dataID++
		out <- fmt.Sprintf("publisher %d is pushing data %d", thisID, dataID)
	}
}
