package arrayalgorithm

import (
	"testing"
)

func TestFindKSmallest(t *testing.T) {
	test := []int {5,4,3,2,1}
	got := FindKSmallest(test,3)
	want := 3
	if want!=got {
		t.Errorf("expected %v got %v\n",want,got)
	}
}
