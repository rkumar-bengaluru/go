package main

import (
	"fmt"
	"strings"
)

func checkCondition(s string) bool {
	if len(s) < 2 {
		return false
	}
	if s[0] == s[1] {
		return false
	}
	for i := 2; i < len(s); i++ {
		if s[i] != s[i%2] {
			return false
		}
	}
	return true
}

func findUniqueString(s string) string {
	var m = make(map[rune]int)
	var res = []byte{}
	for _, k := range s {
		if _, ok := m[k]; !ok {
			m[k] = 1
			res = append(res, byte(k))
		} else {
			m[k] = m[k] + 1
		}
	}
	fmt.Println(m)
	return string(res)
}

// abc
func findPairs(s []byte, l, r int) map[string]string {
	var res = make(map[string]string)
	for i := l; i <= r; i++ {
		for j := i + 1; j <= r; j++ {
			strWithoutPair := strings.ReplaceAll(string(s), string(s[i]), "")
			strWithoutPair = strings.ReplaceAll(strWithoutPair, string(s[j]), "")
			b := string(s[i]) + string(s[j])
			res[b] = strWithoutPair
			//res = append(res, []string{string(s[i]), strWithoutPairstring(s[j])})
		}
	}
	return res
}

func removeEachChar(src string, rep string) string {
	var res string
	res = src
	for _, v := range rep {
		res = strings.ReplaceAll(res, string(v), "")
	}
	return res
}

func alternate(s string) int32 {
	var max int32
	unique := findUniqueString(s)
	pairs := findPairs([]byte(unique), 0, len(unique)-1)
	fmt.Println(pairs)
	for _, p := range pairs {
		// remove chars from s and check condition
		str := removeEachChar(s, p)

		if checkCondition(str) {
			fmt.Println(str)
			if int32(len(str)) > max {
				max = int32(len(str))
			}
		}
	}
	return max
}

func main() {
	// s := []string{"abaacdabd", "beabeefeab","asdcbsdcagfsdbgdfanfghbsfdab"}
	// e := []string{"bdbd", "babab"}
	// fmt.Println(s[0], check(s[0]))
	// fmt.Println(s[1], check(s[1]))
	// fmt.Println(e[0], check(e[0]))
	// fmt.Println(e[1], check(e[1]))
	// fmt.Println(checkWithPair(findUnique(s[0]), s[0]))
	//s := "asdcbsdcagfsdbgdfanfghbsfdab"
	//fmt.Println(alternate("babab"))
	fmt.Println(alternate("beabeefeab"))

}
