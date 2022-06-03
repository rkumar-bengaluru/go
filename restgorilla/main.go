package main

import (
	"github.com/rkumar-bengaluru/go/restgorilla"
)

var a = restgorilla.App{}

func main() {

	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)

	a.Run(":8080")
}
