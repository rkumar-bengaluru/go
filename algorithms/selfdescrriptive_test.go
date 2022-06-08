package algorithms

import (
	"testing"
)

func TestIfDescriptive(t *testing.T) {
	num := 6210001000
	if r := CheckIfSelfDescriptive(num); !r {
		t.Errorf("the num %v is self descriptive\n", num)
	}

	num = 6210001001
	if r := CheckIfSelfDescriptive(num); r {
		t.Errorf("the num %v is not self descriptive\n", num)
	}
}
