package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/consts"
	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/request"
	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/response"
	v "github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/validator"
	"github.com/InsaneProgZ/sale-system-backend/domain/service"
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
	var createProductRequest request.CreateProductRequest

	createProductRequest, err := v.ValidateCreateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}
	log.Println(fmt.Printf("Creating product %+v", createProductRequest))

	product, err := controller.Service.CreateProduct(request.CreateRequestToDomain(createProductRequest))

	if err != nil {
		handler(err, writer)
		return
	}

	response.SetResponse(writer, http.StatusCreated, []consts.Header{
		consts.ContentType,
		consts.AccessControlAllowHeaders,
		consts.AccessControlAllowOrigin,
		consts.AccessControlAllowMethods},
		response.DomainToResponse(product))
}

func (controller *ControllerImpl) FindAllProducts(writer http.ResponseWriter, httpRequest *http.Request) {
	log.Println("Find all products")
	products, err := controller.Service.FindAllProducts()
	if err != nil {
		handler(err, writer)
		return
	}

	response.SetResponse(writer, http.StatusOK, []consts.Header{
		consts.ContentType,
		consts.AccessControlAllowHeaders,
		consts.AccessControlAllowOrigin},
		response.ProductsDomainToProductsResponse(products))
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

	response.SetResponse(writer, http.StatusOK, []consts.Header{consts.ContentType, consts.AccessControlAllowHeaders, consts.AccessControlAllowOrigin}, response.DomainToResponse(product))
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
	updateRequest, err := v.ValidateUpdateRequest(httpRequest.Body, writer)
	if err != nil {
		return
	}
	err = controller.Service.ChangeProductByCode(code, request.UpdateRequestToDomain(updateRequest))
	if err != nil {
		handler(err, writer)
	}
	response.SetResponse(writer, http.StatusOK, []consts.Header{consts.ContentType, consts.AccessControlAllowHeaders, consts.AccessControlAllowOrigin}, nil)
}

func (controller *ControllerImpl) OptionsForBrowsers(writer http.ResponseWriter, httpRequest *http.Request) {
	response.SetResponse(writer, http.StatusOK, []consts.Header{consts.ContentType, consts.AccessControlAllowHeaders, consts.AccessControlAllowOrigin}, nil)
}
