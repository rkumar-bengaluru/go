package sorting

import (
	"reflect"
	"testing"
)

func TestFindMinAndMax(t *testing.T) {
	num := []int{5, 4, 3, 2, 1}
	max, min := FindMinAndMax(num, 0)

	if max != 0 {
		t.Errorf("expected %v got %v\n", 0, max)
	}

	if min != 4 {
		t.Errorf("expected %v got %v\n", 4, min)
	}
}

func TestFindMinAndMax01(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	max, min := FindMinAndMax(num, 0)

	if max != 4 {
		t.Errorf("expected %v got %v\n", 4, max)
	}

	if min != 0 {
		t.Errorf("expected %v got %v\n", 0, min)
	}
}

func TestSelectionSort(t *testing.T) {
	num := []int{5, 4, 3, 2, 1}
	got := SelectionSort(num)
	want := []int{1, 2, 3, 4, 5}
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}
}

func TestSelectionSort02(t *testing.T) {
	num := []int{1, 2, 3, 5, 4}
	got := SelectionSort(num)
	want := []int{1, 2, 3, 4, 5}
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}
}
