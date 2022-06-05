package main

import (
	routes "github.com/rkumar-bengaluru/go/restwgin/nroutes"
)

func main() {
	app := routes.App{}
	app.InitializeRoutes()
	app.Router.Run("localhost:8080")
}
