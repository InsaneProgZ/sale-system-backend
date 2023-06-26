package controller

import (
	"log"
	"net/http"
	"strings"
)

func handler(err error, writer http.ResponseWriter) {
	log.Println(err)

	if strings.Contains(err.Error(), CodeErrors["name_exists"]) {
		errorResponse := GenericErrorResponse{
			Message: "Product already registered!",
		}
		setResponse(writer, http.StatusUnprocessableEntity, []header{contentType}, errorResponse)
		return
	} else if strings.Contains(err.Error(), CodeErrors["id_not_exists"]) {
		errorResponse := GenericErrorResponse{
			Message: "Product id not exists!",
		}
		setResponse(writer, http.StatusUnprocessableEntity, []header{contentType}, errorResponse)
		return
	} else {
		errorResponse := GenericErrorResponse{
			Message: "Internal Server Error, try again in few moments",
		}
		setResponse(writer, http.StatusInternalServerError, []header{contentType}, errorResponse)
		return
	}

}
