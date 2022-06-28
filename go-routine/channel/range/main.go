package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for m := range c {
			fmt.Println("recvd ", m)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	time.Sleep(time.Second)
}
