package sorting

import (
	"math"
)

func FindMinAndMax(n []int, start int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32
	minIdx := 0
	maxIdx := 0
	for i := start; i < len(n); i++ {
		if n[i] < min {
			minIdx = i // 5
			min = n[i]
		}
		if n[i] > max {
			maxIdx = i
			max = n[i]
		}
	}
	return maxIdx, minIdx
}

func SelectionSort(n []int) []int {
	l := len(n) - 1
	for i := 0; i <= l; i++ {
		_, idx := FindMinAndMax(n, i)
		n[i], n[idx] = n[idx], n[i]
	}
	return n
}
