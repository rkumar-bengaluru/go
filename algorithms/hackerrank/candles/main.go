package main

import "fmt"

func distribute(d []int32) {
	var c = make([]int32, len(d))
	c[0] = 1
	var i int32
	for i = 1; i < int32(len(d)); i++ {
		if d[i] > d[i-1] {
			// increase 1 from previous assigned candles
			c[i] = c[i-1] + 1
		} else {
			c[i] = 1
		}
	}
	for i = int32(len(d)) - 2; i >= 0; i-- {
		if d[i] > d[i+1] && c[i] <= c[i+1] {
			c[i] = c[i+1] + 1
		}
	}
	fmt.Println(d)
	fmt.Println(c)
}
func main() {
	// d := [][]int32{
	// 	{4, 6, 4, 5, 6, 2},
	// 	{1, 2, 2},
	// 	{2, 4, 2, 6, 1, 7, 8, 9, 2, 1},
	// 	{2, 4, 3, 5, 2, 6, 4, 5},
	// }
	r := [][]int32{
		{2, 4, 2, 6, 1, 7, 8, 9, 3, 2},
		{1, 2, 1, 2, 1, 2, 3, 4, 1, 2},
	}
	distribute(r[0])
}
