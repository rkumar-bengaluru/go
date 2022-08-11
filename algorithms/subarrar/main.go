package main

import (
	"fmt"
	"strconv"
)

// 1 2 3 4
type DD struct {
	list [][]int
}

var l DD

func allsubarray(d []byte, start, end int32) {
	if end == int32(len(d)) {
		return
	} else if start > end {
		allsubarray(d, 0, end+1)
	} else {
		p := []int{}
		for i := start; i <= end; i++ {
			v, _ := strconv.Atoi(string(d[i]))
			p = append(p, v)
		}
		l.list = append(l.list, p)
		allsubarray(d, start+1, end)
	}
	return
}

func perMutationOfString(d []byte, start, end int32) {
	if start == end {
		fmt.Println(string(d))
	}
	var i int32
	for i = start; i <= end; i++ {
		d[i], d[start] = d[start], d[i]
		perMutationOfString(d, start+1, end)
		d[i], d[start] = d[start], d[i]
	}
}

func dynamic(d []int64) {
	var r = make([]int64, len(d))
	r[0] = d[0]
	for i := 1; i < len(d); i++ {
		if r[0] > 0 {
			r[i] = d[i] + r[i-1]
		} else {
			r[i] = d[i]
		}
	}
	fmt.Println(r)
}
func main() {
	d := []int64{1, 2, 3, 4}
	//perMutationOfString(d, 0, int32(len(d)-1))
	dynamic(d)
	//allsubarray(d, 1, int32(len(d)-1))

}
