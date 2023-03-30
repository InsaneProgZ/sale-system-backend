package handler

import (
	"encoding/json"
	"net/http"
	"sale-system/src/model/request"
	"sale-system/src/service"
)


func Handler(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "POST":
		{
			var productRequest request.Product

			err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
			if err != nil {
				println(err)
				panic(err)
			}

			product := productRequest.ToDomain()

			productResponse := service.CreateProduct(product).ToResponse()

			responseBody, err := json.Marshal(productResponse)
			if err != nil {
				println(err)
				panic(err)
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusCreated)
			writer.Write(responseBody)
		}
		case "GET":
		{
			service.FindAllProducts()
		}
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}
