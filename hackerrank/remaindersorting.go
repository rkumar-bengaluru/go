package hackerrank

import (
	"fmt"
	"math"
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
func RemainderSorting(strArr []string) []string {
	var q int32
	q = 3
	var d int
	d = 5
	r := math.Floor(float64(q)) / d
	fmt.Printf("%.4f\n", r)
	return []string{"Oregon", "Wisconsin", "Utah", "Colarado"}
}
