package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func check(a, b, r int32) bool {
	v := math.Pow(float64(a), float64(b))
	return v == float64(r)
}

// 10 , 2
// 1^2, 2^2, 3^2, 4^2, 5^2

// 100 , 2
// 10^2
var count float64

// {1 , 2 , 3 , 4  }
func powerSum(X float64, N float64) float64 {
	// Write your code here
	if X == 1 {
		return 1
	}
	//		1 2
	n := powerSum(X, N) + powerSum(X-1, N)
	if math.Pow(n, N) == X {
		count++
	}
	return count
}
func sumDigit(n *big.Int) int64 {
	fmt.Println("s", n)

	s := n.String()
	if len(s) == 1 {
		return n.Int64()
	}
	fmt.Println("s", s)
	//r := []rune(s)
	var sum int64
	for _, v := range s {
		i, _ := strconv.ParseInt(string(v), 10, 64)
		sum += i
	}
	fmt.Println("sum", sum, sum)
	return sumDigit(big.NewInt(sum))
}

func superDigit(n string, k int32) int32 {
	// Write your code here
	fmt.Println(n, k)
	if len(n) == 1 {
		i, _ := strconv.ParseInt(n, 10, 64)
		return int32(i)
	}
	i, _ := big.NewInt(1).SetString(n, 10)
	i = i.Mul(i, big.NewInt(int64(k)))
	return int32(sumDigit(i))
}

func main() {
	result := superDigit("861568688536788", 100000)
	fmt.Println(result)
}
