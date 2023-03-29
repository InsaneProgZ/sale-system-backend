package domain

import (
	"sale-system/src/model/response"

	"github.com/google/uuid"
)

type Product struct {
	Name      string
	BuyValue  uint64
	SellValue uint64
	Brand     string
	Code      uuid.UUID
}

func (product Product) ToResponse() response.Product {
	return response.Product{
		Name:      product.Name,
		BuyValue:  product.BuyValue,
		SellValue: product.SellValue,
		Brand:     product.Brand,
		Code:      product.Code,
	}

}
