package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")

	//n := big.NewInt(40)
	start := time.Now()
	res := factorial(25)
	fmt.Println(time.Since(start))
	fmt.Println(res)

}

func factorial(n int64) int64 {
	if n == 1 {
		return 1
	}
	time.Sleep(10 * time.Millisecond)
	return n * factorial(n-1)
}
