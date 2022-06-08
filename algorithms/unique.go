package algorithms

import (
	"fmt"
	"strconv"
)

// unique if all digits in a num is unique
// example: 11 is not unique
func CheckIfUnique(n int) bool {
	utf := strconv.Itoa(n)
	utfs := []byte(utf)
	for pos, v := range utfs {
		i, _ := strconv.Atoi(string(v))
		// check if i is present in the remaining bytes
		fmt.Println(pos, len(utfs))
		rem := utfs[pos : len(utfs)-1]
		fmt.Println(rem)
		for j := 0; j < len(rem)-1; j++ {
			k, _ := strconv.Atoi(string(rem[j]))
			if i == k {
				return false
			}
		}
	}
	fmt.Printf("provided number %v is unique\n", n)
	return true
}
