package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"sale-system/src/model/web_request"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateCreateRequest(body io.ReadCloser, writer http.ResponseWriter) (request web_request.CreateProductRequest, err error) {
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
	var errorResponses []BadRequestResponse

	for _, err := range err.(validator.ValidationErrors) {

		requestType := reflect.TypeOf(request)
		field, _ := requestType.FieldByName(err.Field())

		fieldName := field.Tag.Get("json")
		tag := err.Tag()
		param := err.Param()

		if param != "" {
			param = " " + param
		}

		log.Printf("Field '%s' is %s %s", fieldName, ValidationsMessage[tag], param)

		errorResponses = append(errorResponses, BadRequestResponse{fieldName, ValidationsMessage[tag] + param})
	}
	json.NewEncoder(writer).Encode(errorResponses)
	return
}

func ValidateUpdateRequest(body io.ReadCloser, writer http.ResponseWriter) (request web_request.UpdateProductRequest, err error) {
	err = json.NewDecoder(body).Decode(&request)

	if err != nil {
		setResponse(writer, http.StatusBadRequest, []header{contentType}, nil)
		return
	}

	err = validate.Struct(request)
	if err == nil {
		return
	}

	var errorResponses []BadRequestResponse
	requestType := reflect.TypeOf(request)

	for _, err := range err.(validator.ValidationErrors) {

		field, _ := requestType.FieldByName(err.Field())

		fieldName := field.Tag.Get("json")
		tag := err.Tag()
		param := err.Param()

		if param != "" {
			param += " "
		}

		errorResponses = append(errorResponses, BadRequestResponse{fieldName, ValidationsMessage[tag] + param})
	}
	setResponse(writer, http.StatusBadRequest, []header{contentType}, errorResponses)
	return
}
