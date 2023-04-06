package controller

import (
	"encoding/json"
	"net/http"
	"sale-system/src/model/web_request"
	"sale-system/src/service"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateProduct(writer http.ResponseWriter, httpRequest *http.Request) {

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

func FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {

	products := service.FindAllProducts()
	responseBody, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)

}

func FindProductById(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		panic(err.Error())
	}
	product := service.FindProductById(code)
	responseBody, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)
}
