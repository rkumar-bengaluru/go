package gologconsumer

import (
	"github.com/rkumar-bengaluru/go/golog"
)

func UseGoLog(message string) int {
	l := golog.Default()
	l.Info(message)
	return 0
}
