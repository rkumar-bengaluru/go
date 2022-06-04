package gologconsumer

import (
	"github.com/rkumar-bengaluru/go/golog"
)

func UseGoLog(message string) {
	l := golog.Default()
	l.Info(message)
}
