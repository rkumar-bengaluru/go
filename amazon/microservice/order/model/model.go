package model

type Order struct {
	ID  int `json:"id"`  // order id
	PID int `json:"pid"` // product id
	CID int `json:"cid"` // customer id
}

type IOrder interface {
	CreateOrder(pid, cid int) Order
	GetOrder(id int) Order
}
