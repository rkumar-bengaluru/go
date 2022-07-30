package main

import (
	"fmt"
	"strings"
)

func check(s string) bool {
	r := []rune(s)
	if len(r) <= 1 {
		return false
	}
	first := r[0]
	second := r[1]
	if first == second {
		return false
	}
	for i := 2; i < len(s); i++ {
		if i%2 == 0 {
			if r[0] != r[i] {
				return false
			}
		} else {
			if r[1] != r[i] {
				return false
			}
		}
	}
	return true
}

func findUnique(s string) [][]rune {
	r := []rune(s)
	m := map[rune]rune{}
	res := []rune{}
	pair := [][]rune{}
	for _, v := range r {
		fmt.Println(m[v], v)
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] = 1 + m[v]
		}
	}
	for k, _ := range m {
		res = append(res, k)
	}
	fmt.Println(m)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			fmt.Println("{", string(res[i]), string(res[j]), "}")
			pair = append(pair, []rune{res[i], res[j]})
		}
	}
	fmt.Println(pair)
	return pair
}

func checkWithPair(p [][]rune, s string) int {
	fmt.Println("Length", len(p[0]))
	m := -1
	for i := 0; i < len(p); i++ {
		s := strings.ReplaceAll(s, string(p[i][0]), "")
		s = strings.ReplaceAll(s, string(p[i][1]), "")
		fmt.Println(s)
		//
		if check(s) {
			if len(s) > m {
				m = len(s)
			}
		}
	}
	return m
}
func main() {
	// s := []string{"abaacdabd", "beabeefeab"}
	// e := []string{"bdbd", "babab"}
	// fmt.Println(s[0], check(s[0]))
	// fmt.Println(s[1], check(s[1]))
	// fmt.Println(e[0], check(e[0]))
	// fmt.Println(e[1], check(e[1]))
	// fmt.Println(checkWithPair(findUnique(s[0]), s[0]))
	s := "asdcbsdcagfsdbgdfanfghbsfdab"
	fmt.Println(checkWithPair(findUnique(s), s))
}
