package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"sale-system/src/model/web_request"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateCreateRequest(request web_request.Product, writer http.ResponseWriter) (err error) {

	err = validate.Struct(request)
	if err == nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	var errorResponses []ErrorResponse

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

		errorResponses = append(errorResponses, ErrorResponse{fieldName, ValidationsMessage[tag] + param})
	}
	json.NewEncoder(writer).Encode(errorResponses)
	return
}
