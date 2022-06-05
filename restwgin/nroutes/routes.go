package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rkumar-bengaluru/go/restwgin/model"
)

type App struct {
	Router *gin.Engine
}

func getProductById(c *gin.Context) {
	id := c.Param("id")
	x, _ := strconv.Atoi(id)
	m := model.Product{ID: x}
	m.GetProduct()
	if m.Name == "" {

		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "product not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, m)
}

func (app *App) InitializeRoutes() {
	app.Router = gin.Default()
	app.Router.GET("/product/:id", getProductById)
}
