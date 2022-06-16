package arrayalgorithm

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	test := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

	got := BinarySearch(test, 0, len(test)-1, 23)
	if got != 5 {
		t.Errorf("expected index of %v is %v but got %v\n", 23, 5, got)
	}
}

func TestRightRotate(t *testing.T) {
	got := []int{1, 2, 3, 4, 5}
	want := []int{5, 1, 2, 3, 4}

	Rotate(got, true, 1)
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}

	got = []int{1, 2, 3, 4, 5}
	want = []int{2, 3, 4, 5, 1}

	Rotate(got, false, 1)
	r = reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("expected %v got %v\n", want, got)
	}
}
func TestFindSecondLargest(t *testing.T) {
	num := []int{5, 4, 3, 2, 1}
	expected := []int{5, 4}
	got := FindSecondLargest(num)
	r := reflect.DeepEqual(expected, got)
	if !r {
		t.Errorf("expected %v got %v\n", expected, got)
	}

	num = []int{1, 2, 3, 4, 5}

	got = FindSecondLargest(num)
	r = reflect.DeepEqual(expected, got)
	if !r {
		t.Errorf("expected %v got %v\n", expected, got)
	}
}

func TestFindMaxAndMin(t *testing.T) {
	num := []int{5, 4, 3, 2, 1}
	expected := []int{5, 1}
	got := FindMinAndMax(num)
	r := reflect.DeepEqual(expected, got)
	if !r {
		t.Errorf("expected %v got %v\n", expected, got)
	}

	num = []int{1, 2, 3, 4, 5}

	got = FindMinAndMax(num)
	r = reflect.DeepEqual(expected, got)
	if !r {
		t.Errorf("expected %v got %v\n", expected, got)
	}
}
