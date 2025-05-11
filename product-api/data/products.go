package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Products map[int]*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func UpdateProduct(id int, prod *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	productList[pos] = prod

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	if item, exists := productList[id]; exists {
		return item, id, nil
	}

	return nil, -1, ErrProductNotFound
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList[p.ID] = p
}

func getNextID() int {
	maxID := 0
	for id := range productList {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

var productList = map[int]*Product{
	1: &Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	2: &Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
