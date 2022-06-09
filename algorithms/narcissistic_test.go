package algorithms

import (
	"testing"
)

// Narcissistic Number is a number that is the
// sum of its own digits each raised to the power
// of the number of digits
func TestCheckIfNarcissistic(t *testing.T) {
	num := 153
	if r := CheckIfNarcissistic(num); !r {
		t.Errorf("expected %v to be narcisstic", num)
	}
}
