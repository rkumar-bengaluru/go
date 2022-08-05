package main

import (
	"fmt"
	"strconv"
)

func reverseBits(n []int64) {

	for _, i := range n {
		k := uint32(i)
		s := fmt.Sprintf("%032b", k)
		r := make([]byte, 32)
		var idx int
		for _, s := range s {
			if s == '0' {
				r[idx] = '1'
			} else {
				r[idx] = '0'
			}
			idx++
		}
		i, _ := strconv.ParseInt(string(r), 2, 64)
		fmt.Println(s, string(r), i)
	}

}

func main() {
	fmt.Println("bitflip")
	var d []int64
	d = []int64{2147483647, 1, 0, 4, 123456, 802743475, 35601423}
	reverseBits(d)

}
