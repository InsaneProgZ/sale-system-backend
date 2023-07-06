package service

import (
	"sale-system/model/domain"
	"sale-system/repository"
)

type ProductService interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	FindAllProducts() ([]domain.Product, error)
	FindProductByCode(id int64) (domain.Product, error)
	ChangeProductByCode(id int64, oldProduct *domain.Product) error
}

type ProductServiceImpl struct {
	Repository repository.Database
}

func (productService *ProductServiceImpl) CreateProduct(product *domain.Product) (_ *domain.Product, err error) {

	product.Code, err = productService.Repository.Save(product)
	if err != nil {
		return
	}

	return product, err
}

func (productService *ProductServiceImpl) FindAllProducts() (_ []domain.Product, err error) {
	return productService.Repository.FindAll()
}

func (productService *ProductServiceImpl) FindProductByCode(id int64) (product domain.Product, err error) {
	product, err = productService.Repository.FindByCode(id)
	return
}

func (productService *ProductServiceImpl) ChangeProductByCode(id int64, oldProduct *domain.Product) (err error) {
	err = productService.Repository.ChangeProductByCode(id, *oldProduct)
	return
}
