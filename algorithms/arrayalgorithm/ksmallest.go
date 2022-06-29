package arrayalgorithm

import (
	"math"
)

func removeElement(n []int, r int) []int {
	res := []int{}
	for i := 0; i < len(n); i++ {
		if n[i] != r {
			res = append(res, n[i])
		}
	}
	return res
}

func FindKSmallest(n []int, k int) int {
	smallest := math.MaxInt32
	l := n
	for k > 0 {
		for i := 0; i < len(l); i++ {
			if l[i] < smallest {
				smallest = l[i]
			}
		}
		l = removeElement(l, smallest)

		if k != 1 {
			smallest = math.MaxInt32
		}
		k--
	}

	return smallest
}
