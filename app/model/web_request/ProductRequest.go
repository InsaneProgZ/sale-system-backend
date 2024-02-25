package web_request

import (
	"sale-system/model/domain"
	"time"
)

type CreateProductRequest struct {
	Name      string `json:"name" validate:"required"`
	BuyPrice  uint64 `json:"buy_price" validate:"gt=0"`
	SellPrice uint64 `json:"sell_price" validate:"gt=0"`
	Brand     string `json:"brand" validate:"required"`
}

type UpdateProductRequest struct {
	Name      string `json:"name" validate:"required_without_all=BuyPrice SellPrice Brand"`
	BuyPrice  uint64 `json:"buy_price" validate:"required_without_all=Name SellPrice Brand"`
	SellPrice uint64 `json:"sell_price" validate:"required_without_all=BuyPrice Name Brand"`
	Brand     string `json:"brand" validate:"required_without_all=BuyPrice SellPrice Name"`
}

func (product CreateProductRequest) ToDomain() domain.Product {
	time := time.Now().UTC()
	return domain.Product{
		Name:          product.Name,
		BuyPrice:      product.BuyPrice,
		SellPrice:     product.SellPrice,
		Brand:         product.Brand,
		Creation_date: time,
	}
}

func (product UpdateProductRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:          product.Name,
		BuyPrice:      product.BuyPrice,
		SellPrice:     product.SellPrice,
		Brand:         product.Brand,
		Creation_date: time.Now(),
	}
}
