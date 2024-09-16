package service

import (
	"github.com/InsaneProgZ/sale-system-backend/adapter/output/mysql"
	"github.com/InsaneProgZ/sale-system-backend/domain/model"
)

type ProductService interface {
	CreateProduct(product model.Product) (model.Product, error)
	FindAllProducts() ([]model.Product, error)
	FindProductByCode(id int64) (model.Product, error)
	ChangeProductByCode(id int64, oldProduct model.Product) error
}

type ProductServiceImpl struct {
	Repository mysql.Database
}

func (productService *ProductServiceImpl) CreateProduct(product model.Product) (_ model.Product, err error) {
	
	product.Code, err = productService.Repository.Save(product)
	if err != nil {
		return
	}

	return product, err
}

func (productService *ProductServiceImpl) FindAllProducts() (_ []model.Product, err error) {
	return productService.Repository.FindAll()
}

func (productService *ProductServiceImpl) FindProductByCode(id int64) (product model.Product, err error) {
	product, err = productService.Repository.FindByCode(id)
	return
}

func (productService *ProductServiceImpl) ChangeProductByCode(id int64, oldProduct model.Product) (err error) {
	err = productService.Repository.ChangeProductByCode(id, oldProduct)
	return
}
