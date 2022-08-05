package main

import "fmt"

func caesarCipher(s string, k int32) string {
	var res = make([]byte, len(s))
	var i int32
	// Write your code here
	for _, v := range s {
		var l int32
		if v >= 97 && v <= 122 {
			// (119 + 87) = 206 % 122 = 84
			l = (v + k%26)
			if l > 122 {
				l = 97 + l%122 - 1
			}

			res[i] = byte(l)
		} else if v >= 65 && v <= 90 {
			l = (v + k%26)
			if l > 90 {
				l = 65 + l%90 - 1
			}
			res[i] = byte(l)
		} else {
			res[i] = byte(v)
		}

		i++
	}
	return string(res)
}

func main() {
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(caesarCipher(s, 3))
	fmt.Println(87 % 26)
}
