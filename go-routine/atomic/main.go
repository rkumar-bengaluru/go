package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int32

func main() {
	wg.Add(2)
	go increment("foo")
	go increment("bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
func increment(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {

		<-time.After(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		atomic.AddInt32(&counter, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt32(&counter))
	}
	wg.Done()
}
