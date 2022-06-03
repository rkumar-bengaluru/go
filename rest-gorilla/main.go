package main

import (
	"github.com/rkumar-bengaluru/go/rest-gorilla/app"
)

var a = tmp.App{}

func main() {

	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)

	a.Run(":8080")
}
