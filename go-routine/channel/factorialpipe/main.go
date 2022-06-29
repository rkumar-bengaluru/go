package main

import "fmt"

func main() {
	in := pipe()
	out := factorial(in)
	for res := range out {
		fmt.Println(res)
	}
}

func pipe() <-chan int64 {
	in := make(chan int64)
	go func() {
		for i := 1; i < 11; i++ {
			for j := 1; j < 11; j++ {
				in <- int64(j)
			}
		}
		defer close(in)
	}()
	return in
}
func factorial(ch <-chan int64) chan int64 {
	out := make(chan int64)
	go func() {
		for n := range ch {
			out <- int64(origf(n))
		}
		close(out)
	}()
	return out
}

func origf(n int64) int64 {
	res := int64(1)
	for n >= 1 {
		res *= n
		n--
	}
	return res
}
