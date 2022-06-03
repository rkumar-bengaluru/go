package main

import (
	"github.com/rkumar-bengaluru/go/rest/routes"
)

func main() {

	var a = routes.App{}
	a.Initialize(
		"postgres",
		"postgrespw",
		"sampledb",
	)

	a.Run(":8080")
}
