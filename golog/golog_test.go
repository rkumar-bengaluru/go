package golog_test

import (
	"github.com/rkumar-bengaluru/go/golog"
	"testing"
)

func TestInfo(t *testing.T) {
	l := golog.Default()
	if l == nil {
		t.Fatal("failed in loading golog")
	}
}
