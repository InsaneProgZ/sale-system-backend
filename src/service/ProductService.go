package service

import (
	"sale-system/src/model/domain"
	"sale-system/src/repository"

	"github.com/google/uuid"
)

func CreateProduct(product domain.Product) *domain.Product {
	product.Code, _ = uuid.NewUUID()

	err := repository.Save(product)
	if err != nil {
		println(err)
		panic(err)
	}

	return &product
}
