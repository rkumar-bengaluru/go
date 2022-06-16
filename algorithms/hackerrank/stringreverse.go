package hackerrank

import (
	"fmt"
	"regexp"
)

// remove all spaces & digits from string
// then reverse the string.
func ReverseString(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	n := reg.ReplaceAllString(str, "")
	fmt.Println(n)
	r := []byte(n)
	i := 0
	j := len(n) - 1

	for i <= j/2 {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}
