package web_request

import (
	"sale-system/src/model/domain"
	"time"
)

type Product struct {
	Name      string `json:"name"`
	BuyPrice  uint64 `json:"buy_price"`
	SellPrice uint64 `json:"sell_price"`
	Brand     string `json:"brand"`
}

func (product Product) ToDomain() domain.Product {
	return domain.Product{
		Name:          product.Name,
		BuyPrice:      product.BuyPrice,
		SellPrice:     product.SellPrice,
		Brand:         product.Brand,
		Creation_date: time.Now().UTC(),
	}
}
