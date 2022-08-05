package main

import (
	"fmt"
	"math"
)

func isPow(n int64) int64 {

	//fmt.Println(strconv.FormatInt(int64(n-1), 2))
	var i float64
	for i = 0; i <= float64(n/2)+1; i++ {
		if math.Pow(2, i) > float64(n) {
			i = i - 1
			break
		} else if math.Pow(2, i) == float64(n) {
			return n / 2
		}
	}
	r := int64(float64(n) - math.Pow(2, i))
	if r < 0 {
		panic("something wrong")
	}
	return r
}

func mask(n int64) int64 {
	var f float64
	f = float64(n)
	return int64(1 << int64(math.Floor(math.Log(f/2))))
}
func NumOfSetBits(n int64) int64 {
	var count int64
	count = 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func counterGame(n int64) string {
	var l int
	for n != 1 {
		n = isPow(n)
		l = ^l
		fmt.Println(n, l)
	}
	if l == -1 {
		return "Louise"
	} else {
		return "Richard"
	}
}
func main() {
	var n int64
	// 	1560834904 Richard
	// 1768820483 Richard
	// 1533726144 Louise
	// 1620434450 Richard
	// 1463674015 Louise
	// 6 Richard
	// 132 Louise
	n = 132 //

	fmt.Println("res", counterGame(n))
}
