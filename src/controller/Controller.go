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

type Controller interface {
	CreateProduct(writer http.ResponseWriter, httpRequest *http.Request)
	FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request)
	FindProductById(writer http.ResponseWriter, httpRequest *http.Request)
	OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request)
}

type ControllerImpl struct {
	Service service.ProductService
}

func (controller *ControllerImpl) CreateProduct(writer http.ResponseWriter, httpRequest *http.Request) {

	var productRequest web_request.Product

	err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
	if err != nil {
		panic(err)
	}

	productResponse := controller.Service.CreateProduct(productRequest.ToDomain()).ToResponse()

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

func (controller *ControllerImpl) FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {

	products := controller.Service.FindAllProducts()
	responseBody, err := json.Marshal(domain.ProductsDomainToProductsResponse(products))
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)

}

func (controller *ControllerImpl) FindProductById(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		panic(err.Error())
	}
	product := controller.Service.FindProductById(code)
	responseBody, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseBody)
}

func (controller *ControllerImpl) OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
}
