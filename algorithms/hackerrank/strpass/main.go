package main

import (
	"fmt"
)

func repeatedString(s string, n int64) {
	// Write your code here
	var times int64
	if n%int64(len(s)) == 0 {
		times = n / int64(len(s))
	} else {
		times = n/int64(len(s)) + 1
	}
	var count int64
	for _, v := range s {
		if v == 'a' {
			count++
		}
	}

	var ans, diff, i, diffc int64
	tot := times * int64(len(s)) // tot length
	if tot > n {
		diff = tot - n
	}
	for i = 0; i < int64(len(s))-diff; i++ {
		if s[i] == 'a' {
			diffc++
		}
	}
	if n%int64(len(s)) == 0 {
		ans = times * count
	} else {
		ans = (times - 1) * count
		ans += diffc

	}
	fmt.Println(s, times, count, ans)
}

func main() {
	repeatedString("abcac", 10)
	repeatedString("aba", 10)
	repeatedString("a", 1000000000000)
	repeatedString("ceebbcb", 817723)
	repeatedString("gfcaaaecbg", 547602)
}
