package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	num := []int{5, 4, 3, 2, 1}
	got := MergeSort(num, 0, len(num)-1)
	want := []int{1, 2, 3, 4, 5}
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}
}

func TestMergeSort02(t *testing.T) {
	num := []int{1, 2, 3, 5, 4}
	got := MergeSort(num, 0, len(num)-1)
	want := []int{1, 2, 3, 4, 5}
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}
}
