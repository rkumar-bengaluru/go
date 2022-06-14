package hackerrank

import (
	"reflect"
	"testing"
)

func TestRemainderSorting(t *testing.T) {
	test := []string{"Colarado", "Utah", "Wisconsin", "Oregon"}
	want := []string{"Oregon", "Wisconsin", "Utah", "Colarado"}

	got := RemainderSorting(test)
	r := reflect.DeepEqual(got, want)
	if !r {
		t.Errorf("want %v got %v\n", want, got)
	}
}
