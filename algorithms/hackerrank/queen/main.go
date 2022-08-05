package main

import "fmt"

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	var row, col int32
	points := [][]int32{}

	for row = 1; row < r_q; row++ {
		for col = 1; col < c_q; col++ {
			if row == col {
				points = append(points, []int32{row, col})
			}
		}
	}
	fmt.Println(points)
	return -1
}
func main() {
	fmt.Println(8 % 5)
	queensAttack(8, 1, 4, 4, [][]int32{})
}
