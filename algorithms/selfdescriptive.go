package algorithms

import (
	"strconv"
)

func CheckIfSelfDescriptive(n int) bool {
	str := strconv.Itoa(n)
	b := []byte(str)
	for pos, v := range b {
		val, _ := strconv.Atoi(string(v))
		// find how many of pos is present in the whole number
		// and that should match with val
		count := 0
		for _, v1 := range b {
			toCompare, _ := strconv.Atoi(string(v1))
			if toCompare == pos {
				count++
			}
		}
		if count != val {
			return false
		}
	}
	return true
}
