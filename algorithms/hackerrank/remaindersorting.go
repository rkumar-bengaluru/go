package hackerrank

import (
	"sort"
)

// sort based on following condition
// sort by string lengths modulo 3
// if multiple strings have same modulo then sort by alphabatical order
// Example :-
// strArr := {"Colarado","Utah","Wisconsin","Oregon"}
// there lengths are {8,4,9,6}
// remainders are    {2,1,0,0}
// hence they response would be
// ["Oregon", "Wisconsin", "Utah", "Colarado"]

type RemainderSortingResponse struct {
	idx int
	mod int
	str string
	s chan
}

type RemainderSortingList struct {
	list []RemainderSortingResponse
}

func (r *RemainderSortingList) Add(l RemainderSortingResponse) {
	r.list = append(r.list, l)
}

func (r *RemainderSortingList) Len() int {
	return len(r.list)
}

func (r *RemainderSortingList) Less(i, j int) bool {
	if r.list[i].mod == r.list[j].mod {
		return r.list[i].str < r.list[j].str
	}
	return r.list[i].mod < r.list[j].mod
}

func (r *RemainderSortingList) Swap(i, j int) {
	r.list[i], r.list[j] = r.list[j], r.list[i]
}

func RemainderSorting(strArr []string) []string {
	// find the modulo based on the string lengths
	size := len(strArr) - 1
	r := RemainderSortingList{}
	for i := 0; i <= size; i++ {
		m := len(strArr[i]) % 3
		r.Add(RemainderSortingResponse{idx: i, mod: m, str: strArr[i]})
	}
	sort.Sort(&r)
	res := []string{}
	for i := 0; i < r.Len(); i++ {
		res = append(res, r.list[i].str)
	}
	return res
}
