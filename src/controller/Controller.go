package controller

import (
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
	FindProductByCode(writer http.ResponseWriter, httpRequest *http.Request)
	ChangeProductByCode(writer http.ResponseWriter, httpRequest *http.Request)
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

func (controller *ControllerImpl) FindProductByCode(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		handler(err, writer)
		return
	}

	product, err := controller.Service.FindProductByCode(code)
	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusOK, []header{contentType}, product.ToResponse())
}

func (controller *ControllerImpl) ChangeProductByCode(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		handler(err, writer)
		return
	}
	request, err := ValidateCreateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}
	err = controller.Service.ChangeProductByCode(code, request.ToDomain())
	if err != nil {
		handler(err, writer)
	}
}

func (controller *ControllerImpl) OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	writer.Header().Set("Access-Control-Allow-Headers", "content-type")
}
