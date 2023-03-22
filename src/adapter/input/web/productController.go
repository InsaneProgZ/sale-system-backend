package web

import (
	"encoding/json"
	"net/http"
	request "sale-system/src/adapter/input/web/request"
	response "sale-system/src/adapter/input/web/response"
	"github.com/google/uuid"
)

func Application() {
	http.HandleFunc("/product", productController)
	http.ListenAndServe("localhost:8081", nil)
}

func productController(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "POST":
		var productRequest request.Product
		err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
		if err != nil {
			http.Error(writer, "Error trying create product", http.StatusBadRequest)
			return
		}
		uuid, _ := uuid.NewUUID()
		productResponse := response.Product{Code: uuid, Name: productRequest.Name, Value: productRequest.Value, Brand: productRequest.Brand}
		responseBody, err := json.Marshal(productResponse)

		if err != nil {
			http.Error(writer, "Error to encode responsebody", http.StatusBadRequest)

			println(err)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		writer.Write(responseBody)
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}
