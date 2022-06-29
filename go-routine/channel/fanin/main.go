package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	in := fanIn(boring("rupak"), boring("raj"))
	for r := range in {
		fmt.Println(r)
	}
}
func boring(s string) chan string {
	out := make(chan string)
	go func() {
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
	}()
	return out
}

func fanIn(i1, i2 chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case result := <-i1:
				out <- result
			case result := <-i2:
				out <- result
			}
		}
	}()

	return out
}
