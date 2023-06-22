package controller

import (
	"encoding/json"
	"net/http"
	"sale-system/src/model/domain"
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
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(responseBody)
	println(string(responseBody))
}

func FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {

	products := service.FindAllProducts()
	responseBody, err := json.Marshal(domain.ProductsDomainToProductsResponse(products))
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
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

func OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
}
