package router

import (
	"net/http"
	"sale-system/src/handler"

	"github.com/gorilla/mux"
)

func Router() {
	mux := mux.NewRouter()

	mux.HandleFunc("/products", handler.HandlerProduct)

	http.ListenAndServe("localhost:8081", mux)

}
