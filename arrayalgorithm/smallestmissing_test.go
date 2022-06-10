package arrayalgorithm

import (
	"testing"
)

func TestSmallestMissing(t *testing.T) {
	test := []int{0, 1, 2, 6, 9}
	n := 5
	m := 10
	got := FindSmallestMissing(test, n, m)
	want := 3

	if got != 3 {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{4, 5, 10, 11}
	n = 4
	m = 12
	got = FindSmallestMissing(test, n, m)
	want = 0

	if got != 0 {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{0, 1, 2, 3}
	n = 4
	m = 5
	got = FindSmallestMissing(test, n, m)
	want = 4

	if got != want {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{0, 1, 2, 3, 4, 5, 6, 7, 10}
	n = 9
	m = 11
	got = FindSmallestMissing(test, n, m)
	want = 8

	if got != want {
		t.Errorf("expecting %v got %v\n", want, got)
	}
}

func TestFindSmallestMissingUsingBinarySearch(t *testing.T) {
	test := []int{0, 1, 2, 6, 9}
	n := 5
	m := 10
	got := FindSmallestMissingUsingBinarySearch(test, n, m)
	want := 3

	if got != 3 {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{4, 5, 10, 11}
	n = 4
	m = 12
	got = FindSmallestMissingUsingBinarySearch(test, n, m)
	want = 0

	if got != 0 {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{0, 1, 2, 3}
	n = 4
	m = 5
	got = FindSmallestMissingUsingBinarySearch(test, n, m)
	want = 4

	if got != want {
		t.Errorf("expecting %v got %v\n", want, got)
	}

	test = []int{0, 1, 2, 3, 4, 5, 6, 7, 10}
	n = 9
	m = 11
	got = FindSmallestMissingUsingBinarySearch(test, n, m)
	want = 8

	if got != want {
		t.Errorf("expecting %v got %v\n", want, got)
	}
}
