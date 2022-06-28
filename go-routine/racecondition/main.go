package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int32

var mutex sync.Mutex

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
		mutex.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
