package web_request

import (
	"sale-system/src/model/domain"
	"time"
)

type Product struct {
	Name      string `json:"name"`
	BuyValue  uint64 `json:"buy_value"`
	SellValue uint64 `json:"sell_value"`
	Brand     string `json:"brand"`
}

func (product Product) ToDomain() domain.Product {
	return domain.Product{
		Name:          product.Name,
		BuyValue:      product.BuyValue,
		SellValue:     product.SellValue,
		Brand:         product.Brand,
		Creation_date: time.Now().UTC(),
	}
}
