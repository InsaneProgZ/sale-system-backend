package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Handler(err error, writer http.ResponseWriter) {
	
	log.Println(err)

	if strings.Contains(err.Error(), CodeErrors["name_exists"]) {
		errorResponse := GenericErrorResponse{
			Message: "Product already registered!",
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(writer).Encode(errorResponse)
		return
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Internal Server Error, try again in few moments")
		return
	}
}
