package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

const productDBName = "products"

var productColumns = []string{"category", "product_name", "price"}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
	records, err := u.db.Load(productDBName)
	if err != nil {
		records = [][]string{productColumns}
		if err := u.db.Save(productDBName, records); err != nil {
			return nil, err
		}
	}

	result := make([]Product, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		product := Product{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
		}
		result = append(result, product)
	}

	return result, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	result, _ := u.LoadOrCreate()
	return result, nil // TODO: replace this
}
