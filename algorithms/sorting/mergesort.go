package sorting

import (
	"math"
	"sort"
)

func MergeSort(n []int, l, r int) []int {
	if l < r {
		// math floor
		m := l + int(math.Floor((float64((r - l) / 2))))
		MergeSort(n, l, m)
		MergeSort(n, m+1, r)
		merge(n, l, m, r)
	}
	sort.Slice()
	return n
}

func merge(n []int, l, m, r int) {
	i := l
	j := m + 1

	for i <= m && j <= r {
		if n[i] > n[j] {
			n[i], n[j] = n[j], n[i]
			i++
			j++
			for k := j; k <= r; k++ {
				if n[k-1] > n[k] {
					n[k-1], n[k] = n[k], n[k-1]
				}
			}
			j = m + 1
		} else {
			i++
			j--
		}
	}
}
