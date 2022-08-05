package main

import "fmt"

func xorSeq(l, r int64) {
	var i, res int64
	var d = make([]int64, r+1)
	d[0] = 0
	for i = 1; i <= r; i++ {
		d[i] = d[i-1] ^ i
	}
	fmt.Println(d)
	res = d[l]
	l++
	for l <= r {
		res = res ^ d[l]
		l++
	}
	fmt.Println(res)
}

func main() {
	xorSeq(4, 6)
}
