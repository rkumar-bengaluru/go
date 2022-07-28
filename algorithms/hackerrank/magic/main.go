package main

import (
	"fmt"
	"math"
)

var minCost int64

func formMagicSquare(d [][]int) int64 {
	var digits []int
	digits = []int{8, 3, 4, 1, 5, 9, 6, 7, 2}
	var square [][]int
	square = [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	minCost = math.MaxInt
	dfs(digits, 0, square, d)
	return minCost
}

func dfs(digits []int, cidx int, t [][]int, s [][]int) {
	if cidx > len(digits) {
		return
	}
	var row, col, index int
	var cost int64
	for row = 0; row < len(t); row++ {
		for col = 0; col < len(t); col++ {
			t[row][col] = digits[index]
			v := t[row][col] - s[row][col]
			cost += int64(math.Abs(float64(v)))
			index++
		}
	}
	//fmt.Println("cost", cost)
	if isMagicSquare(t, 15) {
		fmt.Println(t, cost)
		fmt.Println(s)
		if cost < minCost {
			minCost = cost
		}
	}

	for i := cidx; i < len(digits); i++ {
		digits[i], digits[cidx] = digits[cidx], digits[i]
		dfs(digits, cidx+1, t, s)
		digits[i], digits[cidx] = digits[cidx], digits[i]
	}
}

func isMagicSquare(m [][]int, sum int) bool {

	var row, col int

	// check if rem rows sum is same
	for row = 0; row < len(m); row++ {
		var s int
		for col = 0; col < len(m); col++ {
			s += m[row][col]
		}
		//fmt.Println(s)
		if s != sum {
			return false
		}
	}
	// check if rem col sum is same
	for col = 0; col < len(m); col++ {
		var s int
		for row = 0; row < len(m); row++ {
			s += m[col][row]
		}
		//fmt.Println(s)
		if s != sum {
			return false
		}
	}
	// check if diag 1 is same
	var s int
	for row = 0; row < len(m); row++ {
		s += m[row][row]
	}
	//fmt.Println(s)
	if s != sum {
		return false
	}
	// check if diag 1 is same
	s = 0
	l := len(m) - 1

	for row = 0; row < len(m); row++ {
		s += m[l-row][row]
	}
	//fmt.Println(s)
	if s != sum {
		return false
	}
	return true
}

func main() {
	// d := [][]int{
	// 	{4, 8, 2}, //-2,-1,4
	// 	{4, 5, 7}, //4,-1,4
	// 	{6, 1, 6}, // 1,0,3
	// }
	// d := [][]int{
	// 	{2, 9, 8}, //-1,0,-3
	// 	{4, 2, 7}, //-2,4,0
	// 	{5, 6, 7}, // -1,-3,1
	// }
	//fmt.Println(formMagicSquare(d))
	var s []int32
	s = make([]int32, 0, 3)
	s[0] = 1
	fmt.Println(len(s), cap(s))

	//fmt.Println(r)
}
