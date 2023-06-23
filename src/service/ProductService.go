package service

import (
	"sale-system/src/model/domain"
	"sale-system/src/repository"
)

type ProductService interface {
	CreateProduct(product domain.Product) domain.Product
	FindAllProducts() []domain.Product
	FindProductById(id int64) domain.Product
}

type ProductServiceImpl struct {
	Database repository.Database
}

func (productService *ProductServiceImpl) CreateProduct(product domain.Product) domain.Product {

	product.Code = productService.Database.Save(product)
	return product
}

func (productService *ProductServiceImpl) FindAllProducts() []domain.Product {
	return productService.Database.FindAll()
}

func (productService *ProductServiceImpl) FindProductById(id int64) domain.Product {
	return productService.Database.FindById(id)
}
