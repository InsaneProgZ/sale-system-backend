package controller

import (
	"encoding/json"
	"log"
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

	var request web_request.Product

	request, err := ValidateCreateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}

	product, err := controller.Service.CreateProduct(request.ToDomain())

	if err != nil {
		Handler(err, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(product.ToResponse())
	if err != nil {
		log.Println(err)
	}
}

func (controller *ControllerImpl) FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {

	products, err := controller.Service.FindAllProducts()
	if err != nil {
		Handler(err, writer)
		return
	}

	responseBody, err := json.Marshal(domain.ProductsDomainToProductsResponse(products))
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}
	product, _ := controller.Service.FindProductById(code)
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
