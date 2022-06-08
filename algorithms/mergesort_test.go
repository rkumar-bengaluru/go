package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	test := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	MergeSort(&test)
	if r := reflect.DeepEqual(test, expected); !r {
		t.Errorf("expected %v got %v\n", expected, test)
	}
}
