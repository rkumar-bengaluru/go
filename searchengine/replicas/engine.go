package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web          = fakeSearch("Web")
	WebReplica   = fakeSearch("Web Replica")
	Image        = fakeSearch("Image")
	ImageReplica = fakeSearch("Image Replica")
	Video        = fakeSearch("Video")
	VideoReplica = fakeSearch("Video Replica")
)

func main() {
	SearchMain("Golang")
}

func GetFromFirstReplica(query string, replicas ...FakeSearch) string {
	c := make(chan string)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
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
	go func() { resultChannel <- GetFromFirstReplica(query, Web, WebReplica, Web) }()
	go func() { resultChannel <- GetFromFirstReplica(query, Image, ImageReplica, ImageReplica) }()
	go func() { resultChannel <- GetFromFirstReplica(query, Video, VideoReplica, VideoReplica) }()
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
