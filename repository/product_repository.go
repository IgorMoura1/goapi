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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	fmt.Println("Product created with ID:", id)

	return id, nil
}