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

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
    var product model.Product
    query := "SELECT id, product_name, price FROM product WHERE id = $1"
    err := pr.connection.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
    return product, err
}

func (pr *ProductRepository) GetProductByName(name string) ([]model.Product, error) {
    var products []model.Product
    query := "SELECT id, product_name, price FROM product WHERE product_name ILIKE $1"
    rows, err := pr.connection.Query(query, "%"+name+"%")
    if err != nil {
        return products, err
    }
    defer rows.Close()
    for rows.Next() {
        var product model.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
            return products, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
    query := "UPDATE product SET product_name = $1, price = $2 WHERE id = $3"
    _, err := pr.connection.Exec(query, product.Name, product.Price, product.ID)
    return err
}