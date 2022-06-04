package main

import (
	"github.com/rkumar-bengaluru/go/golog"
	"github.com/rkumar-bengaluru/go/gologconsumer"
)

func main() {
	l := golog.Default()
	l.Info("logging from main...")
	gologconsumer.UseGoLog("using golog from gologconsumer package...")
}
