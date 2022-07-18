package model

import "testing"

func TestCreateOrder(t *testing.T) {
	var inmem IOrder
	inmem = &InMemOrder{}
	pid := 2
	cid := 1
	o := inmem.CreateOrder(pid, cid)
	if o.ID != 3 {
		t.Errorf("expected %v got %v", 3, o.ID)
	}
}
