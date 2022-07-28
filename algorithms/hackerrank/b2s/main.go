package main

import (
	"fmt"
	"sort"
)

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	f := []int32{}
	for i := a[len(a)-1]; i <= b[len(b)-1]; i++ {
		if i%a[len(a)-1] == 0 && i <= b[0] {
			k := true
			for _, j := range a {
				fmt.Println(i, j, i%j)
				if i%j != 0 {
					k = false
					break
				}
			}
			if k {
				f = append(f, i)
			}
		}
	}
	r := []int32{}
	for _, i := range f {
		k := true
		for _, j := range b {
			if j%i != 0 {
				k = false
				break
			}
		}
		if k {
			r = append(r, i)
		}
	}
	
	fmt.Println(f)
	fmt.Println(r)
	return -1

}

func main() {
	//fmt.Println(getTotalX([]int32{3, 4}, []int32{24, 48}))
	fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}))
}
