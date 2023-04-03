package handler

import (
	"encoding/json"
	"net/http"
	"sale-system/src/model/web_request"
	"sale-system/src/service"
)
func HandlerProduct(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "POST":
		{
			var productRequest web_request.Product

			err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
			if err != nil {
				panic(err)
			}

			productResponse := service.CreateProduct(productRequest.ToDomain()).ToResponse()

			responseBody, err := json.Marshal(productResponse)
			if err != nil {
				panic(err)
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusCreated)
			writer.Write(responseBody)
		}
	case "GET":
		{
			products := service.FindAllProducts()
			responseBody, err := json.Marshal(products)
			if err != nil {
				panic(err)
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(responseBody)
		}
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}

func HandlerProductId(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "GET":
		{
			println(httpRequest.RequestURI)
			service.FindAllProducts()
		}
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}
