package main

import (
	"fmt"
	"math"
)

func check(i int32, y []int32) bool {
	for _, v := range y {
		if v == i {
			return true
		}
	}
	return false
}

func weightedUniformStrings(s string, queries []int32) []string {
	// Write your code here
	// a = 97 || 1 || 97 -97 + 1 = 1
	// b = 98 || 2 || 98-97 + 1 = 2
	// y = 121 || 25 || 121-97 + 1 = 25
	r := []rune(s)
	var w = make([]int32, len(r))
	var res = make([]string, len(queries))
	var prev rune
	var length int32
	prev = ' '
	// abccddde
	// 123648125
	for i := 0; i < len(r); i++ {
		v := int32(r[i]) - 97 + 1 // {1,2,4,3,6,9,4,8,12,16}
		w[i] = v                  // 1,2
		if prev == r[i] {
			length++
			w[i] = v * length
		} else {
			prev = r[i]
			length = 1
		}

	}
	var i int32
	for _, q := range queries {
		if check(q, w) {
			res[i] = "Yes"
		} else {
			res[i] = "No"
		}
		i++
	}
	fmt.Println()
	return res
}

func main() {
	s1 := "abccddde"
	q1 := []int32{1, 3, 12, 5, 9, 10}
	//q1 := []int32{9, 7, 8, 12, 5}
	v := weightedUniformStrings(s1, q1)
	fmt.Println(v)
	fmt.Println(math.Abs(float64(110 - 111)))

}
