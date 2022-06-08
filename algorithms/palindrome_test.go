package algorithms

import "testing"

func TestStringPalindrom(t *testing.T) {
	test := "radar"
	if r := CheckIfStringPalindrome(test); !r {
		t.Errorf("was expecting %v to be palindrome", test)
	}
}

func TestIntPalindrom(t *testing.T) {
	var test int = 12621
	if r := CheckIfIntPalindrome(test); !r {
		t.Errorf("was expecting %v to be palindrome", test)
	}
}
