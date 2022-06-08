package algorithms

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if r, _ := Reverse("rupak"); r != "kapur" {
		t.Errorf("expecting kapur got %s\n", r)
	}
}

func TestReverse01(t *testing.T) {
	if r, _ := Reverse("rupakc"); r != "ckapur" {
		t.Errorf("expecting kkapur got %s\n", r)
	}
}
