package arrayalgorithm

func FindDuplicates(n []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(n); i++ {
		if m[n[i]] != 0 {
			m[n[i]] = m[n[i]] + 1
		} else {
			m[n[i]] = 1
		}
	}
	return m
}
