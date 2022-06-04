package gologconsumer_test

import (
	"github.com/rkumar-bengaluru/go/gologconsumer"
	"testing"
)

func TestUseGoLog(t *testing.T) {

	if r := gologconsumer.UseGoLog("something"); r != 0 {
		t.Fatal("failed Test")
	}
}
