package algorithms

import (
	"math"
)

func MergeSort(n *[]int) {
	start := 0
	end := len(*n) - 1
	sort(n, start, end)
}

func sort(input *[]int, l, r int) {
	if l < r {
		// math floor
		var m float64 = float64(l) + math.Floor((float64(r)-float64(l))/2)
		var middle = int(m)
		sort(input, l, middle)
		sort(input, middle+1, r)
		merge(input, l, middle, r)
	}
}

func merge(input *[]int, l, m, r int) {
	var i = l
	var j = m + 1

	for i <= m && j <= r {
		if (*input)[i] > (*input)[j] {
			// swap
			(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
			i++
			j++
			for k := j; k <= r; k++ {
				if (*input)[k-1] > (*input)[k] {
					(*input)[k-1], (*input)[k] = (*input)[k], (*input)[k-1]
				}
			}
			j = m + 1
		} else {
			i++
			j++
		}
	}
}
