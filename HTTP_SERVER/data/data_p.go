package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"myname" validate:"required"`
	Description string  `json:"description" validate:"gt=0"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
}
type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (P *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(P)
}
func (P *Product) ToDATA(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(P)
}
func GetProduct(id int) *Product {
	_, pos, err := findProduct(id)
	if err != nil {
		fmt.Errorf("Product not found")
	}
	return ProductList[pos]
}
func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	ProductList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
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
