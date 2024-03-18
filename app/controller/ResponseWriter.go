package controller

import (
	"encoding/json"
	"net/http"
)

var Headers = map[header]string{
	"Content-Type":                 "application/json",
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Headers": "*",
	"Access-Control-Allow-Methods": "*",
}

type header string

const (
	contentType               header = "Content-Type"
	AccessControlAllowOrigin  header = "Access-Control-Allow-Origin"
	AccessControlAllowHeaders header = "Access-Control-Allow-Headers"
	AccessControlAllowMethods header = "Access-Control-Allow-Methods"
)

func getHeader(header header) (key string, value string) {
	return string(header), Headers[header]
}

func setResponse(writer http.ResponseWriter, statusCode int, headers []header, body interface{}) {
	for _, header := range headers {
		writer.Header().Set(getHeader(header))
	}
	writer.WriteHeader(statusCode)

	if body == nil {
		body = ""
	}
	json.NewEncoder(writer).Encode(body)
}
