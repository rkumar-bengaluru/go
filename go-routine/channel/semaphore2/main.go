package main

import "fmt"

func main() {
	c := make(chan int)
	b := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()

	go func() {
		for m := range c {
			fmt.Println(m)
		}
		b <- true
	}()
	go func() {
		for m := range c {
			fmt.Println(m)
		}
		b <- true
	}()
	<-b
	<-b
}
