/*
   Problem Statement
   Given an array with n elements, can you sort this array in ascending order using just one of
   the following operations? You can perform only one of the following operations:
   1. Swap two elements.
   2. Reverse one sub-segment
   Input Format
   The first line contains a single integer n, which indicates the size of the array.
   The next line contains n integers seperated by spaces.
   n
   d1 d2 ... dn
   Constraints
   2 <= n <= 100000
   0 <= di <= 1000000
   All di are distinct.
   Output Format
       If the array is already sorted, output 'yes' in the first line.
       You do not need to output anything else.
       If you can sort this array using one single operation (from the two permitted operations):
       a. If you can sort the array by swap dl and dr, output "swap l r" in the second line.
	      l and r are the indices of the elements to be swapped, assuming that the array is
		  indexed from 1 to n.
       b. Else if it is possible to sort the array by reversing the segment d[l...r], output
	      "reverse l r" in the second line. l and r are the indices of the first and last elements
		  of the subsequence to be reversed, assuming that the array is indexed from 1 to n.
          d[l...r] represents the sub-sequence of the array, beginning at index l and ending at
		  index r; inclusive of both.
       If an array can be sorted by either swapping or reversing, stick to the swap based method.
       If you cannot sort the array in either of the above ways, output "no" in the first line.
   Sample Input #1
   2
   4 2
   2 4
   ---
   2 2
   Sample Output #1
   yes
   swap 1 2
   Sample Input #2
   3
   3 1 2
   Sample Output #2
   no
   Sample Input #3
   6
   1 5 4 3 2 6 || 6 2 3 4 5 1
   Sample Output #3
   yes
   reverse 2 5
   Explanation
   For #1: You can both swap(1, 2) and reverse(1, 2), but if you can sort the array but swap, output swap only.
   For #2: It is impossible to sort by one single operation (among those permitted).
   For #3, You can reverse the sub-array d[2...5] = "5 4 3 2" then the array become sorted.
   1 5 4 3 2 6
   1 2 3 4 5 6
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func checkIfSorted(d []int64) bool {
	for i := 1; i < len(d); i++ {
		if d[i] < d[i-1] {
			return false
		}
	}
	return true
}
func asort(d []int64, n int64, sorted []int64) string {
	i := int64(0)
	l := n
	j := l - 1
	br := 0
	for i < j && br < 2 {
		if d[i] == sorted[i] {
			i++
		} else {
			br++
		}

		if d[j] == sorted[j] {
			j--
		} else {
			br++
		}
	}
	fmt.Printf("i = %v, j = %v br = %v\n", i, j, br)
	if i == j {
		return fmt.Sprint("yes")
	}

	ii := i + 1
	ij := j + 1
	times := 0

	br = 0

	for i < j && br < 1 {
		br = 0

		if d[i] == sorted[j] && d[j] == sorted[i] {
			times++
			i++
			j--
		} else {
			br++
		}
	}
	fmt.Printf("i = %v, j = %v br = %v times=%v\n", i, j, br, times)
	if i >= j {
		if times == 1 {
			return fmt.Sprintf("yes\nswap %v %v", ii, ij)
		}
		return fmt.Sprintf("yes\nreverse %v %v", ii, ij)
	}
	return ""
}

func main() {
	r := bufio.NewReader(os.Stdin)
	l, _, _ := r.ReadLine()
	n, _ := strconv.ParseInt(strings.TrimSpace(string(l)), 10, 64)
	fmt.Println("no of elements", n)
	nl, _, _ := r.ReadLine()

	arr := strings.Split(strings.TrimSpace(string(nl)), " ")
	var narr []int64
	for _, a := range arr {
		n, _ := strconv.ParseInt(a, 10, 64)
		narr = append(narr, n)
	}
	c := make([]int64, n)
	copy(c, narr)

	sort.SliceStable(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	fmt.Println("copy", c)
	fmt.Println(narr)
	asort(narr, n, c)
}
