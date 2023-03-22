package service

import (
	"sale-system/src/application/domain"
	"github.com/google/uuid"
)

func CreateProduct(product domain.Product) domain.Product{
	product.Code, _ = uuid.NewUUID()
	return product
}