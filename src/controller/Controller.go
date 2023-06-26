package controller

import (
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
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusCreated, []header{contentType}, product.ToResponse())
}

func (controller *ControllerImpl) FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {

	products, err := controller.Service.FindAllProducts()
	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusOK, []header{contentType}, domain.ProductsDomainToProductsResponse(products))
}

func (controller *ControllerImpl) FindProductById(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		log.Println(err)
		handler(err, writer)
		return
	}

	product, err := controller.Service.FindProductById(code)
	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusOK, []header{contentType}, product)
}

func (controller *ControllerImpl) OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
}
