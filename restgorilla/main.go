package main

import (
	"github.com/rkumar-bengaluru/go/restgorilla/app"
)

var a = app.App{}

func main() {

	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)

	a.Run(":8080")
}
