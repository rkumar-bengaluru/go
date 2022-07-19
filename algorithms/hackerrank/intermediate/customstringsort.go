package main

import (
	"fmt"
	"sort"
)

type CustomStringList struct {
	l []CustomString
}
type CustomString struct {
	data string
	len  int
	even int
}

func (l *CustomStringList) Len() int {
	return len(l.l)
}

func (l *CustomStringList) Less(i, j int) bool {
	if l.l[i].even != 0 && l.l[j].even != 0 {
		// odd
		if l.l[i].len == l.l[j].len {
			return l.l[i].len > l.l[j].len
		} else if l.l[i].even != 0 && l.l[j].even != 0 {
			return l.l[i].len < l.l[j].len
		}
	} else {
		// even
		if l.l[i].len == l.l[j].len {
			return l.l[i].len < l.l[j].len
		} else if l.l[i].even != 0 && l.l[j].even != 0 {
			return l.l[i].len < l.l[j].len
		}
	}
	return true
}

func (l *CustomStringList) Swap(i, j int) {
	l.l[i], l.l[j] = l.l[j], l.l[i]
}

func main() {
	t := []string{"abc", "ab", "abcde", "a", "abcd"}
	cs := CustomStringList{}
	for _, s := range t {
		if len(s)%2 != 0 {
			cs.l = append(cs.l, CustomString{data: s, len: len(s), even: (len(s) % 2)})
		}
	}
	sort.Sort(&cs)
	res := []string{}
	for _, r := range cs.l {
		res = append(res, r.data)
	}
	cs = CustomStringList{}
	for _, s := range t {
		if len(s)%2 == 0 {
			cs.l = append(cs.l, CustomString{data: s, len: len(s), even: (len(s) % 2)})
		}
	}
	sort.Sort(&cs)
	for _, r := range cs.l {
		res = append(res, r.data)
	}
	fmt.Println(res)
}
