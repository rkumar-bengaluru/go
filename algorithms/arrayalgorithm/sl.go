package main

import (
	"fmt"
	"math"
)

func slargest(d []int) []int {
	largest := math.MinInt
	slargest := largest

	for _, i := range d {
		if largest < i {
			slargest, largest = largest, i
		} else if slargest < i {
			slargest = i
		}
	}
	return []int{largest, slargest}
}
func main() {
	var d = []int{5, 4, 3, 2, 1}
	fmt.Println(slargest(d))
}
