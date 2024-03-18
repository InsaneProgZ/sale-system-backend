package web_request

import (
	"sale-system/model/domain"
	"time"
)

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Price uint64 `json:"price" validate:"gt=0"`
	Brand string `json:"brand" validate:"required"`
}

type UpdateProductRequest struct {
	Name  string `json:"name" validate:"required_without_all=Price Brand"`
	Price uint64 `json:"price" validate:"required_without_all=Name Brand"`
	Brand string `json:"brand" validate:"required_without_all=Price Name"`
}

func (product CreateProductRequest) ToDomain() domain.Product {
	time := time.Now().UTC()
	return domain.Product{
		Name:          product.Name,
		Price:         product.Price,
		Brand:         product.Brand,
		Creation_date: time,
	}
}

func (product UpdateProductRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:          product.Name,
		Price:         product.Price,
		Brand:         product.Brand,
		Creation_date: time.Now(),
	}
}
