package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func working() {
	c := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(c)
		for i := 0; i < 9; i++ {
			c <- i
		}
	}()
	go func() {
		defer wg.Done()
		for r := range c {
			fmt.Printf("recvd on channel %v\n", r)
		}
	}()
	wg.Wait()
}

func working2() {
	oddChan := make(chan int)
	evenChan := make(chan int)

	wg.Add(3)
	go odd(oddChan)
	go even(evenChan)
	go func() {
		defer wg.Done()
		defer close(oddChan)
		defer close(evenChan)

		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				evenChan <- i
			} else {
				oddChan <- i
			}
		}
	}()

	wg.Wait()

}

func main() {
	oddChan := make(chan int)
	evenChan := make(chan int)
	test := []int{5, 4, 3, 2, 1}
	wg.Add(3)
	go Server1()
	go odd(oddChan)
	go even(evenChan)
	for i := 0; i < len(test); i++ {
		serverChan <- in{i: test[i], oddChan: oddChan, evenChan: evenChan}
	}
	close(serverChan)
	wg.Wait()
}

type in struct {
	i        int
	oddChan  chan int
	evenChan chan int
}

var serverChan = make(chan in)

func Server1() {
	var input in
	defer wg.Done()
	for input = range serverChan {
		if input.i%2 == 0 {
			input.evenChan <- input.i
		} else {
			input.oddChan <- input.i
		}

	}

	defer close(input.oddChan)
	defer close(input.evenChan)
}

func Server() {

	var input1 in = <-serverChan
	defer wg.Done()

	defer close(input1.oddChan)
	defer close(input1.evenChan)
}

func odd(ch <-chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("ODD :", v)

	}
}

func even(ch <-chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("EVEN:", v)

	}
}
