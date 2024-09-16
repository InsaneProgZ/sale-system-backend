package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/consts"
	"github.com/InsaneProgZ/sale-system-backend/domain/model"
)

type Product struct {
	Code         int64     `json:"code"`
	Name         string    `json:"name"`
	Price        uint64    `json:"price"`
	Brand        string    `json:"brand"`
	CreationDate time.Time `json:"creation_date"`
}

func SetResponse(writer http.ResponseWriter, statusCode int, Headers []consts.Header, body interface{}) {
	for _, Header := range Headers {
		writer.Header().Set(consts.GetHeader(Header))
	}
	writer.WriteHeader(statusCode)

	if body == nil {
		body = ""
	}
	json.NewEncoder(writer).Encode(body)
}

func DomainToResponse(product model.Product) Product {
	return Product{
		Code:         product.Code,
		Name:         product.Name,
		Brand:        product.Brand,
		Price:        product.Price,
		CreationDate: product.Creation_date.In(time.Local),
	}
}

func ProductsDomainToProductsResponse(products []model.Product) []Product {
	var responses []Product

	for _, product := range products {
		a := Product{
			Code:         product.Code,
			Name:         product.Name,
			Brand:        product.Brand,
			Price:        product.Price,
			CreationDate: product.Creation_date.In(time.Local),
		}
		responses = append(responses, a)
	}
	return responses

}
