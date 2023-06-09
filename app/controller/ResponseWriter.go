package controller

import (
	"encoding/json"
	"net/http"
	"reflect"
)

var Headers = map[header]string{
	"Content-Type": "application/json",
}

type header string

const (
	contentType header = "Content-Type"
)

func getHeader(header header) (key string, value string) {
	return string(header), Headers[header]
}

func setResponse(writer http.ResponseWriter, statusCode int, headers []header, body interface{}) {
	for _, header := range headers {
		writer.Header().Set(getHeader(header))
	}
	writer.WriteHeader(statusCode)
	if !reflect.ValueOf(body).IsNil() {
		json.NewEncoder(writer).Encode(body)
	}
}
