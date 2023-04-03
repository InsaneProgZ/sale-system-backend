package domain

import (
	"sale-system/src/model/web_response"
	"time"
)

type Product struct {
	Code          int64
	Name          string
	Brand         string
	BuyValue      uint64
	SellValue     uint64
	Creation_date time.Time
}

func (product Product) ToResponse() web_response.Product {
	return web_response.Product{
		Code:         product.Code,
		Name:         product.Name,
		Brand:        product.Brand,
		BuyValue:     product.BuyValue,
		SellValue:    product.SellValue,
		CreationDate: product.Creation_date.In(time.Local),
	}
}
