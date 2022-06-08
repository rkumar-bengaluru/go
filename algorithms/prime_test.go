package algorithms

import (
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	test01 := []int{2, 3, 5, 7, 11}       // true
	test02 := []int{0, 1, 4, 6, 8, 9, 10} // false

	for i := 0; i < len(test01)-1; i++ {
		n := test01[i]
		r := CheckIfPrime(n)
		if !r {
			t.Errorf("was expecting %d to be prime\n", n)
		}
	}

	for i := 0; i < len(test02)-1; i++ {
		r := CheckIfPrime(test02[i])
		if r {
			t.Errorf("was expecting %d to be not prime\n", test02[i])
		}
	}
}

func TestAllPrime(t *testing.T) {
	test := 12
	expected := []int{2, 3, 5, 7, 11}
	res := GiveAllPrimes(test)
	if r := reflect.DeepEqual(res, expected); !r {
		t.Errorf("expected response to be %v but got %v\n", expected, res)
	}
}
