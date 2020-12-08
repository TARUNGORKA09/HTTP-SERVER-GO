package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	DeletedOn   string
	UpdatedOn   string
}

func GetProduct() []*Product {
	return ProductList
}

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}