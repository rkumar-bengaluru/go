package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rkumar-bengaluru/go/logger/console"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewUserHandler,
	),
	fx.Invoke(Initialize),
)

type UserHandlerInterface interface {
	GetProduct(id string)
}
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type UserHandler struct {
	log    *console.ConsoleLogger
	users  map[string]User
	router *mux.Router
}

func (c *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	c.log.Info(fmt.Sprintf("recvd request for GetUser %v \n", id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	p := c.users[""+id+""]
	re, _ := json.Marshal(p)
	w.Write(re)
}

func NewUserHandler(log *console.ConsoleLogger, router *mux.Router) *UserHandler {
	return &UserHandler{log: log, users: make(map[string]User), router: router}
}

func Initialize(c *UserHandler) {
	c.users["1"] = User{Id: "1", Name: "Rupak"}
	c.users["2"] = User{Id: "2", Name: "Aryan"}

	c.router.HandleFunc("/user/{id}", c.GetUser).Methods("GET")
}
