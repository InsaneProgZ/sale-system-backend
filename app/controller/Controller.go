package controller

import (
	"fmt"
	"log"
	"net/http"
	"sale-system/model/domain"
	"sale-system/model/web_request"
	"sale-system/service"
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
	var request web_request.CreateProductRequest

	request, err := ValidateCreateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}
	log.Println(fmt.Printf("Creating product %+v", request))

	product, err := controller.Service.CreateProduct(request.ToDomain())

	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusCreated, []header{contentType, AccessControlAllowHeaders, AccessControlAllowOrigin, AccessControlAllowMethods}, product.ToResponse())
}

func (controller *ControllerImpl) FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {
	log.Println("Find all products")
	products, err := controller.Service.FindAllProducts()
	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusOK, []header{contentType, AccessControlAllowHeaders, AccessControlAllowOrigin}, domain.ProductsDomainToProductsResponse(products))
}

func (controller *ControllerImpl) FindProductByCode(writer http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	numberOfBits := 64
	numberBase := 10

	code, err := strconv.ParseInt(vars["code"], numberBase, numberOfBits)	
	if err != nil {
		handler(err, writer)
		return
	}
	
	log.Println(fmt.Printf("Creating product %d", code))
	
	product, err := controller.Service.FindProductByCode(code)
	if err != nil {
		handler(err, writer)
		return
	}

	setResponse(writer, http.StatusOK, []header{contentType, AccessControlAllowHeaders, AccessControlAllowOrigin}, product.ToResponse())
}

func (controller *ControllerImpl) ChangeProductByCode(writer http.ResponseWriter, httpRequest *http.Request) {
	contentTypeRequest := httpRequest.Header.Get("Content-Type")
	if contentTypeRequest != "application/json" {
		http.Error(writer, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	vars := mux.Vars(httpRequest)
	code, err := strconv.ParseInt(vars["code"], 10, 64)
	if err != nil {
		handler(err, writer)
		return
	}
	request, err := ValidateUpdateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}
	err = controller.Service.ChangeProductByCode(code, request.ToDomain())
	if err != nil {
		handler(err, writer)
	}
	setResponse(writer, http.StatusOK, []header{contentType, AccessControlAllowHeaders, AccessControlAllowOrigin}, nil)
}

func (controller *ControllerImpl) OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	setResponse(writer, http.StatusOK, []header{contentType, AccessControlAllowHeaders, AccessControlAllowOrigin}, nil)
}
