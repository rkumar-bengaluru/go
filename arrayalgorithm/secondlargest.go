package arrayalgorithm

import (
	"math"
)

func BinarySearch(n []int, l, r, k int) int {
	if l <= r {
		m := l + int(math.Floor(float64((r-l)/2)))
		if k == n[m] {
			return m
		}
		if k > n[m] {
			return BinarySearch(n, m+1, r, k)
		} else {
			return BinarySearch(n, l, m-1, k)
		}
	}
	return -1
}

func Rotate(n []int, right bool, times int) {

	if right {
		for i := 0; i < times; i++ {
			l := len(n) - 1
			r := n[l]
			for l > 0 {
				n[l] = n[l-1]
				l--
			}
			n[0] = r
		}
	}
	// {1,2,3,4,5} -- {2,3,4,5,1}
	if !right {
		for i := 0; i < times; i++ {
			l := len(n) - 1
			r := n[0]
			for i := 0; i < l; i++ {
				n[i] = n[i+1]
			}
			n[l] = r
		}
	}
}

func FindSecondLargest(n []int) []int {
	var largest int = 0
	var second int = 0

	for i := 0; i < len(n); i++ {
		if n[i] > largest {
			second, largest = largest, n[i]
		} else if second < n[i] {
			second = n[i]
		}

	}
	res := []int{largest, second}
	return res
}

func FindMinAndMax(n []int) []int {
	var max int = 0
	var min int = math.MaxInt32
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i] // 5
		}

		if n[i] < min {
			min = n[i] // 5
		}
	}
	return []int{max, min}
}
