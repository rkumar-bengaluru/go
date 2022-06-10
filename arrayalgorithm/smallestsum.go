package arrayalgorithm

import (
	"fmt"
	"sort"
)

type Pair struct {
	num1 int
	num2 int
	sum  int
}

type Pairs struct {
	p map[int]Pair
}

func New() *Pairs {
	return &Pairs{make(map[int]Pair)}
}
func (a Pairs) Len() int {
	return len(a.p)
}
func (a Pairs) Less(i, j int) bool {
	return a.p[i].sum < a.p[j].sum
}
func (a Pairs) Swap(i, j int) {
	a.p[i], a.p[j] = a.p[j], a.p[i]
}
func FindKSmallestSum(n []int, m []int, k int) int {
	r := []int{}
	pairs := New()
	z := 0
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(m); j++ {
			sum := (n[i] + m[j])
			r = append(r, sum)
			pairs.p[z] = Pair{i, j, sum}
			z++
		}
	}
	sort.Sort(pairs)
	fmt.Printf("( %v, %v) sum %v\n", n[pairs.p[0].num1], m[pairs.p[0].num2], pairs.p[0].sum)
	fmt.Printf("( %v, %v) sum %v\n", n[pairs.p[1].num1], m[pairs.p[1].num2], pairs.p[1].sum)
	fmt.Printf("( %v, %v) sum %v\n", n[pairs.p[2].num1], m[pairs.p[2].num2], pairs.p[2].sum)
	return 3
}
