package service

import (
	"log"
	"sale-system/src/model/domain"
	"sale-system/src/repository"
)

type ProductService interface {
	CreateProduct(product domain.Product) (domain.Product, error)
	FindAllProducts() ([]domain.Product, error)
	FindProductById(id int64) (domain.Product, error)
}

type ProductServiceImpl struct {
	Database repository.Database
}

func (productService *ProductServiceImpl) CreateProduct(product domain.Product) (_ domain.Product, err error) {

	product.Code, err = productService.Database.Save(product)
	if err != nil {
		log.Println(err)
		return
	}

	return product, err
}

func (productService *ProductServiceImpl) FindAllProducts() (_ []domain.Product, err error) {
	return productService.Database.FindAll()
}

func (productService *ProductServiceImpl) FindProductById(id int64) (_ domain.Product, err error) {
	return productService.Database.FindById(id)
}
