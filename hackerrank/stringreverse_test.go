package hackerrank

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	test := " ab 12cd e34fg "
	want := "gfedcba"
	got := ReverseString(test)

	if got != want {
		t.Errorf("want %v got %v\n", want, got)
	}
}
