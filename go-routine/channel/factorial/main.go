package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	fmt.Println(a) // 100-digit Fibonacci number
}

func factorial01(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * factorial01(n-1)
}

func factorialBig(n *big.Int) *big.Int {
	var res *big.Int = big.NewInt(1)
	one := big.NewInt(1)
	//z := big.NewInt(0)
	fmt.Println("one.Cmp(one) ", one.Cmp(one))
	fmt.Println("n.Cmp(one) ", n.Cmp(one))
	// fmt.Println("n.Sub(one) ", n.Sub(n, one))
	// fmt.Println("n.Sub(one) ", n.Sub(n, one))
	fmt.Println("n.Sub(one) ", n.Mul(big.NewInt(2), big.NewInt(2)))
	r := 2
	for r >= 0 {
		res = n.Mul(n, res)
		n = n.Sub(n, one)
		r--
		fmt.Println(r, n, res)
	}
	return res
}

func factorial(n int64) chan int64 {
	c := make(chan int64)
	go func() {
		var res int64 = 1
		for n >= 1 {
			res *= n
			n--
		}
		c <- res
		close(c)
	}()

	return c
}
