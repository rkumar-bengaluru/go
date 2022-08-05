package main

import "fmt"

// a b c
func permute(s []rune, l, r int) {
	if l == r {
		fmt.Println(string(s))
		return
	}
	for i := l; i <= r; i++ { //   i= 1, l = 0
		s[i], s[l] = s[l], s[i] // b a c
		permute(s, l+1, r)      //
		s[l], s[i] = s[i], s[l]
	}
}

func main() {
	str := "abc"

	permute([]rune(str), 0, len(str)-1)
}
