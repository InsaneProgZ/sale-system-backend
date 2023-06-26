package domain

import (
	"sale-system/src/model/web_response"
	"time"
)

type Product struct {
	Code          int64
	Name          string
	Brand         string
	BuyPrice      uint64
	SellPrice     uint64
	Creation_date *time.Time
}

func (product Product) ToResponse() web_response.Product {
	return web_response.Product{
		Code:         product.Code,
		Name:         product.Name,
		Brand:        product.Brand,
		BuyPrice:     product.BuyPrice,
		SellPrice:    product.SellPrice,
		CreationDate: product.Creation_date.In(time.Local),
	}
}

func ProductsDomainToProductsResponse(products []Product) []web_response.Product {
	var responses []web_response.Product

	for _, product := range products {
		a := web_response.Product{
			Code:         product.Code,
			Name:         product.Name,
			Brand:        product.Brand,
			BuyPrice:     product.BuyPrice,
			SellPrice:    product.SellPrice,
			CreationDate: product.Creation_date.In(time.Local),
		}
		responses = append(responses, a)
	}
	return responses

}