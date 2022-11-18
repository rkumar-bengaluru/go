package main

import (
	"fmt"
	"strconv"
)

func convertoint(s string) []int64 {
	r := []rune(s)
	var ints = make([]int64, len(r))
	for idx, v := range r {
		i, err := strconv.ParseInt(string(byte(v)), 10, 64)
		if err != nil {
			panic("cannot parse string to int")
		}
		ints[idx] = i
	}
	return ints
}

func getNextValidInt(s []rune) {
	i, _ := strconv.ParseInt(string(byte(s[0])), 10, 64)
	if i == 0 {
		fmt.Println("NO")
		return
	}
	var start int
	res := []int64{}

	for start < len(s) {

		fdigits := []byte{}
		for k := 0; k <= start; k++ {
			fdigits = append(fdigits, byte(s[k]))
			m, _ := strconv.ParseInt(string(fdigits), 10, 64)

			res = append(res, m)
		}
		sdigits := []byte{}
		for j := start + 1; j < len(s); j++ {
			sdigits = append(sdigits, byte(s[j]))
			n, _ := strconv.ParseInt(string(sdigits), 10, 64)
			if n == 0 {
				break
			}

			if n-(res[start]) == 1 {
				res = append(res, n)
			}
		}

		start++
		fmt.Println(res)
	}

	fmt.Println(string(s), "YES")
}

// 1 2 3 4
// 9 10 11 - 9 10 11
// 99100 - 99 100
// 101103 - 10 - 11
// 1 3
func beautifulstring(s string) {
	r := []rune(s)
	getNextValidInt(r)
}
func main() {
	s := "1234"
	beautifulstring(s)
	// beautifulstring("91011")
	// beautifulstring("99100")

	//beautifulstring("101103")
}
