package arrayalgorithm

func FindSmallestMissingUsingBinarySearch(n []int, size, max int) int {
	// if 0 is missing at the first idx then return 0
	if n[0] != 0 {
		return 0
	}
	for i := 0; i < max; i++ {
		if BinarySearch(n, 0, len(n)-1, i) == -1 {
			return i
		}
	}
	return -1
}

func FindSmallestMissing(n []int, size, max int) int {

	// if 0 is missing at the first idx then return 0
	if n[0] != 0 {
		return 0
	}

	for i := 0; i <= size-2; i++ {
		// check if the adjacent difference is
		// greater than 1 , then missing number
		// is value at ith index + 1
		if (n[i+1] - n[i]) > 1 {
			return n[i] + 1
		}
	}
	// finally if the value at last index
	//
	return (max - 1)
}
