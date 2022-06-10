package arrayalgorithm

import (
	"testing"
)

func TestFindKSmallestSum(t *testing.T) {
	n := []int{1, 4, 11}
	m := []int{2, 8, 6}
	got := FindKSmallestSum(n, m, 3)
	want := 3

	if got != want {
		t.Errorf("was expecting %v but got %v\n", want, got)
	}

}
