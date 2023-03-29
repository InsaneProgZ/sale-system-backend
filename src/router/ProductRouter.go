package router

import (
	"net/http"
	"sale-system/src/handler"
)

func Router(){
	mux := http.NewServeMux()

	mux.Handle("/products", http.HandlerFunc(handler.Handler))

	http.ListenAndServe("localhost:8081", mux)

}