package repository

import (
	"database/sql"
	"fmt"
	"goapi/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
		query := "SELECT id, product_name, price FROM product"
		rows, err := pr.connection.Query(query)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		var products []model.Product
		var productObj model.Product
		
		for rows.Next() {
		err = rows.Scan(
			&productObj.ID, 
			&productObj.Name, 
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		products = append(products, productObj)
	}

	rows.Close()
	return products, nil
}