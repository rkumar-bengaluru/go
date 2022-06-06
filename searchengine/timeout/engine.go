package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

func main() {
	SearchMain("Golang")
}

func SearchMain(query string) {
	start := time.Now()
	results := SearchEngine("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

func SearchEngine(query string) (results []string) {
	resultChannel := make(chan string)
	quit := time.After(80 * time.Millisecond)
	go func() { resultChannel <- Web(query) }()
	go func() { resultChannel <- Image(query) }()
	go func() { resultChannel <- Video(query) }()
	for i := 0; i < 3; i++ {
		select {
		case result := <-resultChannel:
			results = append(results, result)
		case <-quit:
			fmt.Println("timeout on search")
			return
		}
	}
	return
}

type FakeSearch func(query string) string

func fakeSearch(kind string) FakeSearch {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%v result for %q\n", kind, query)
	}
}
