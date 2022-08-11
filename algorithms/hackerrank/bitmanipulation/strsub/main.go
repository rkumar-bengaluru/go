package main

import (
	"fmt"
)

func hackerrankInString(s string) string {
	// Write your code here
	var nextChar []byte
	var j int
	nextChar = []byte{104, 97, 99, 107, 101, 114, 114, 97, 110, 107}
	for i := 0; i < len(s); i++ {
		if j < len(nextChar) {
			if s[i] == nextChar[j] {
				j++
			}
		}

	}
	fmt.Println(j)
	if j != 10 {
		return "NO"
	}
	return "YES"
}

// 2 2 4 4 5 8 -- 2 2 3 6 - 1 4 - 3
// 0 0 2 2 3 6 - 6
// 0 0 0 0 1 4 - 4
// 0 0 0 0 0 3 - 2
//
// 6 4 2 1
func main() {

	fmt.Println("111%7", 1011%7)
	s := "hhhackkerbanker"
	fmt.Println(hackerrankInString(s))

}
