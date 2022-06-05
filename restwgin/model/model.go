package model

import "errors"

type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var all = map[int]Product{
	1: {1, "apple", 8.99},
	2: {2, "mango", 5.99},
}

func (p *Product) GetProduct() error {
	for k, val := range all {
		if k == p.ID {
			p.Name = val.Name
			p.Price = val.Price
			return nil
		}
	}
	return errors.New("could not find product")
}
