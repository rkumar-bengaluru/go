package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	// d1 := []int32{
	// 	4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2,
	// 	98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3,
	// 	4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2,
	// 	98, 4, 2, 1, 1, 9, 2, 4,
	// }
	// d := []int32{
	// 	7, 12, 13, 19, 17, 7, 3, 18, 9, 18, 13, 12, 3, 13, 7, 9, 18, 9, 18, 9, 13, 18, 13,
	// 	13, 18, 18, 17, 17, 13, 3, 12, 13, 19, 17, 19, 12, 18, 13, 7, 3, 3, 12, 7, 13, 7,
	// 	3, 17, 9, 13, 13, 13, 12, 18, 18, 9, 7, 19, 17, 13, 18, 19, 9, 18, 18, 18, 19, 17,
	// 	7, 12, 3, 13, 19, 12, 3, 9, 17, 13, 19, 12, 18, 13, 18, 18, 18, 17, 13, 3, 18,
	// 	19, 7, 12, 9, 18, 3, 13, 13, 9, 7,
	// }
	// d3 := []int32{
	// 	4, 2, 3, 4, 4, 9, 2, 4,
	// }

	// d4 := []int32{
	// 	66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	// 	// 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66 66
	// }
	//fmt.Println(pickingNumbers(d4))
	s := "300.01"
	i, _ := strconv.ParseFloat(s, 64)
	r := i * 300.00 / 600.00
	r = math.Round(r*100) / 100
	fmt.Println(i, r)

}

func checkCondition(d []int32, n int32) bool {
	for _, r := range d {
		if n-r > 1 {
			return false
		}
	}
	return true
}
func pickingNumbers(a []int32) int32 {
	// Write your code here
	sort.Slice(a, func(i, j int) bool {
		return a[i] <= a[j]
	})
	fmt.Println(a)
	first := a[0]
	var l, max int32
	max = math.MinInt32
	l = 1
	s := []int32{first}

	for i := 1; i < len(a); i++ {
		if a[i] == first || first+1 == a[i] {
			l++
			s = append(s, a[i])
		} else {
			if l > max {
				max = l
			}
			first = a[i]
			s = []int32{first}
			l = 1
		}

	}
	if l > max {
		max = l
	}
	return max
}
