package request

import (
	"sale-system/src/application/domain"
	"github.com/google/uuid"
)

type Product struct {
	Name      string `json:"name"`
	BuyValue  uint64 `json:"buy_value"`
	SellValue uint64 `json:"sell_value"`
	Brand     string `json:"brand"`
}

func (product Product) ToDomain() domain.Product {
	return domain.Product{
		Name:      product.Name,
		BuyValue:  product.BuyValue,
		SellValue: product.SellValue,
		Brand:     product.Brand,
		Code:      uuid.UUID{},
	}
}
