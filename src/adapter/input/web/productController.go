package web

import (
	"encoding/json"
	"net/http"
	"sale-system/src/adapter/input/web/request"
	"sale-system/src/application/handler"
	"sale-system/src/application/service"
)

func Application() {
	http.HandleFunc("/products", productController)
	http.ListenAndServe("localhost:8081", nil)
}

func productController(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "POST":
		{
			var productRequest request.Product

			err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
			handler.ErrorHandler(err)

			product := productRequest.ToDomain()

			productResponse := service.CreateProduct(product).ToResponse()

			responseBody, err := json.Marshal(productResponse)
			handler.ErrorHandler(err)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusCreated)
			writer.Write(responseBody)
		}
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}
