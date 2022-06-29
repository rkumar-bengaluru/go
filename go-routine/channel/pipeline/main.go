package main

import "fmt"

func main() {
	// c := gen(2, 3)
	// r := square(c)

	// fmt.Println(<-r)
	// fmt.Println(<-r)
	// Set up the pipeline and consume the output.
	for n := range square(square(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
func gen(nums ...int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close((out))
	}()
	return out
}

func square(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
