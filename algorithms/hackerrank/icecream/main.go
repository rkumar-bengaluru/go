package main

import "fmt"

func whatFlavors(cost []int32, money int32) {
	var i int32
	var m = make(map[int32]int32)
	for i = 0; i < int32(len(cost)); i++ {
		d := money - cost[i]
		// check if d is present.
		if m[d] == 0 {
			// not found the combo
			if m[cost[i]] == 0 {
				m[cost[i]] = i + 1
			}
		} else {
			// found the combo
			fmt.Println(m[d], i+1)
		}

	}
}
func main() {
	whatFlavors([]int32{2, 2, 4, 3}, 4)
}
