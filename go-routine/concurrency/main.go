package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func foo() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("foo ->", i)
	}
}

func bar() {
	defer wg.Done()
	for i := 0; i < 45; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("bar ->", i)
	}
}
