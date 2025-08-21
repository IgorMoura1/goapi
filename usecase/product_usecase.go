package usecase

import (
    "goapi/model"
    "goapi/repository"
)

type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productID, err := pu.productRepository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID

	return product, nil
}