package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/consts"
	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller/response"
)

func handler(err error, writer http.ResponseWriter) {
	log.Println(err)

	if strings.Contains(err.Error(), response.CodeErrors["name_exists"]) {
		errorResponse := response.GenericErrorResponse{
			Message: "Product already registered!",
		}
		response.SetResponse(writer, http.StatusUnprocessableEntity, []consts.Header{consts.ContentType}, errorResponse)
		return
	} else if strings.Contains(err.Error(), response.CodeErrors["id_not_exists"]) {
		errorResponse := response.GenericErrorResponse{
			Message: "Product id not exists!",
		}
		response.SetResponse(writer, http.StatusUnprocessableEntity, []consts.Header{consts.ContentType}, errorResponse)
		return
	} else {
		errorResponse := response.GenericErrorResponse{
			Message: "Internal Server Error, try again in few moments",
		}
		response.SetResponse(writer, http.StatusInternalServerError, []consts.Header{consts.ContentType}, errorResponse)
		return
	}

}