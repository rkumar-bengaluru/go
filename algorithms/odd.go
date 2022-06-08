package algorithms

func CheckIfAllOdd(n []int) bool {
	for i := 0; i <= len(n)-1; i++ {
		if r := CheckIfOdd(n[i]); !r {
			return false
		}
	}
	return true
}	

func CheckIfOdd(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	// a number is odd if the remainder is not 0 after dividing with 2
	if n%2 != 0 {
		return true
	}
	return false
}
