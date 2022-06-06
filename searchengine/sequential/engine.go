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
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

type FakeSearch func(query string) string

func fakeSearch(kind string) FakeSearch {
	return func(query string) string {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return fmt.Sprintf("%v result for %q\n", kind, query)
	}
}
