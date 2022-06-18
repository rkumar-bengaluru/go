package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductCatalog interface {
	GetProduct(id int) (*Product, error)
	UpdateProduct(p Product) (*Product, error)
	DeleteProduct(id int) (*Product, error)
	CreateProduct(p Product) (*Product, error)
	GetAllProducts() ([]Product, error)
}
