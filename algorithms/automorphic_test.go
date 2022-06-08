package algorithms

import (
	"testing"
)

func TestAutoMorphic(t *testing.T) {
	num := 25

	if r := CheckIfAutomorphic(num); !r {
		t.Errorf("expected %v to be automorphic\n", num)
	}

	num = 49
	if r := CheckIfAutomorphic(num); r {
		t.Errorf("expected %v to be not automorphic\n", num)
	}
}
