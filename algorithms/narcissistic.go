package algorithms

import (
	"math"
	"strconv"
)

// Narcissistic Number is a number that is the
// sum of its own digits each raised to the power
// of the number of digits
func CheckIfNarcissistic(n int) bool {
	// find all digits.
	str := strconv.Itoa(n)
	utf := []byte(str)
	var res int
	for _, v := range utf {
		i, _ := strconv.Atoi(string(v))
		res += int(math.Pow(float64(i), float64(len(utf))))
	}
	if res == n {
		return true
	}
	return false
}
