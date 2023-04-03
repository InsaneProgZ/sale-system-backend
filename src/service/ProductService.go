package service

import (
	"sale-system/src/model/domain"
	"sale-system/src/repository"
)

func CreateProduct(product domain.Product) domain.Product {

	product.Code = repository.Save(product)

	return product
}

func FindAllProducts() []domain.Product{
	return repository.FindAll()
}
