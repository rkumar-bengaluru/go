package algorithms

import (
	"strconv"
)

func CheckIfStringPalindrome(str string) bool {
	b := []rune(str)
	m := (len(b) - 1) / 2
	for i, j := 0, (len(b) - 1); i <= m; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}

	}
	return true
}

func CheckIfIntPalindrome(n int) bool {
	r := strconv.Itoa(n)
	return CheckIfStringPalindrome(r)
}
