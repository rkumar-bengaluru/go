package main

import "fmt"

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

func main() {
	i := 10
	i++

	s := "hhhackkerbanker"
	fmt.Println(hackerrankInString(s))

}
