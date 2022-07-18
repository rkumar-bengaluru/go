package model

var orders = []Order{
	Order{1, 1, 1},
	Order{2, 2, 1},
}

type InMemOrder struct{}

func (mem *InMemOrder) CreateOrder(pid, cid int) Order {
	neworder := Order{3, pid, cid}
	orders = append(orders, neworder)
	return neworder
}

func (mem *InMemOrder) GetOrder(id int) Order {
	for _, o := range orders {
		if o.ID == id {
			return o
		}
	}
	return Order{}
}
