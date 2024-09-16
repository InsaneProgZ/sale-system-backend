package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/consts"
	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/request"
	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/response"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateCreateRequest(body io.ReadCloser, writer http.ResponseWriter) (request request.CreateProductRequest, err error) {
	err = json.NewDecoder(body).Decode(&request)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate.Struct(request)
	if err == nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorResponses []response.BadRequestResponse

	for _, err := range err.(validator.ValidationErrors) {

		requestType := reflect.TypeOf(request)
		field, _ := requestType.FieldByName(err.Field())

		fieldName := field.Tag.Get("json")
		tag := err.Tag()
		param := err.Param()

		if param != "" {
			param = " " + param
		}

		log.Printf("Field '%s' is %s %s", fieldName, response.ValidationsMessage[tag], param)

		errorResponses = append(errorResponses, response.BadRequestResponse{Field: fieldName, Message: response.ValidationsMessage[tag] + param})
	}
	json.NewEncoder(writer).Encode(errorResponses)
	return
}

func ValidateUpdateRequest(body io.ReadCloser, writer http.ResponseWriter) (request request.UpdateProductRequest, err error) {
	err = json.NewDecoder(body).Decode(&request)

	if err != nil {
		response.SetResponse(writer, http.StatusBadRequest, []consts.Header{consts.ContentType}, nil)
		return
	}

	err = validate.Struct(request)
	if err == nil {
		return
	}

	var errorResponses []response.BadRequestResponse
	requestType := reflect.TypeOf(request)

	for _, err := range err.(validator.ValidationErrors) {

		field, _ := requestType.FieldByName(err.Field())

		fieldName := field.Tag.Get("json")
		tag := err.Tag()
		param := err.Param()

		if param != "" {
			param += " "
		}

		errorResponses = append(errorResponses, response.BadRequestResponse{Field: fieldName, Message: response.ValidationsMessage[tag] + param})
	}
	response.SetResponse(writer, http.StatusBadRequest, []consts.Header{consts.ContentType, consts.AccessControlAllowHeaders, consts.AccessControlAllowOrigin}, errorResponses)
	return
}
