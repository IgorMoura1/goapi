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