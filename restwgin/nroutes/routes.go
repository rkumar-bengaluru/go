package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rkumar-bengaluru/go/restwgin/model"
)

type App struct {
	router *gin.Engine
}

func getProductById(c *gin.Context) {
	id := c.Param("id")
	m := model.Product{ID: strconv.Atoi(id)}
	if m.Name != "" {
		c.IndentedJSON(http.StatusNotFound, json.Marshal(`{ "msg" : "Product Not Found"`))
	}

	c.IndentedJSON(http.StatusOK, m)
}

func (app App) InitializeRoutes() {
	app.router = gin.Default()
	app.router.GET("/product/{id}", getProductById)
}
