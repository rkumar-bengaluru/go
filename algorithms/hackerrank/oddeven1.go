package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
)

/*
 * Complete the 'Server' function below and missing types and global variables.
 *
 * The function is void.
 */
type in struct {
	input int32
	even  chan int32
	odd   chan int32
}

var serverChan = make(chan in)

func Server() {
	fmt.Println("inside server")
	// wait for something to be available at serverChan
	var i in
	for i = range serverChan {
		fmt.Println(i)
		e := i.input % 2
		fmt.Printf("remainder for %d is %v\n", i.input, e)
		if e == 0 {
			i.even <- i.input
		} else {
			i.odd <- i.input
		}

	}
	defer close(i.even)
	defer close(i.odd)
}

func main() {

	var arr []int32
	for i := 0; i < 10; i++ {
		arrItem := int32(rand.Intn(100))
		arr = append(arr, arrItem)
	}
	fmt.Println(arr)
	oddChan := make(chan int32)
	evenChan := make(chan int32)
	fmt.Println("--------------")
	fmt.Println("ssss")
	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	fmt.Println("ssss")
	go func() {
		fmt.Println("waiting for odd")
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	go Server()
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], evenChan, oddChan}
	}
	close(serverChan)

	wg.Wait()

	for _, resultItem := range odds {
		fmt.Printf("recv odd %d\n", resultItem)
	}

	for _, resultItem := range evens {
		fmt.Printf("recv even %d\n", resultItem)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
