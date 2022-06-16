package arrayalgorithm

import (
	"reflect"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	test := []int{4, 4, 1, 1, 9, 9, 8, 8, 8}
	got := FindDuplicates(test)
	want := map[int]int{1: 2, 4: 2, 8: 3, 9: 2}

	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("want %v got %v\n", want, got)
	}
}
