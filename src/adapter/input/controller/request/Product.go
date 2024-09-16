package request

import (
	"time"

	"github.com/InsaneProgZ/sale-system-backend/domain/model"
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

func CreateRequestToDomain(product CreateProductRequest) model.Product {
	time := time.Now().UTC()
	return model.Product{
		Name:          product.Name,
		Price:         product.Price,
		Brand:         product.Brand,
		Creation_date: time,
	}
}

func UpdateRequestToDomain(product UpdateProductRequest) model.Product {
	return model.Product{
		Name:          product.Name,
		Price:         product.Price,
		Brand:         product.Brand,
		Creation_date: time.Now(),
	}
}
