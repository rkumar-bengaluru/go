package algorithms

import (
	"testing"
)

func TestAnagrams(t *testing.T) {
	s1 := "danger"
	s2 := "garden"
	got := checkIfAnagrams(s1, s2)
	want := true

	if got != want {
		t.Errorf("got %v want %v for s1 %s and %s2", got, want, s1, s2)
	}
}
