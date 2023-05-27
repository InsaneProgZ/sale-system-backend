package main

import (
	"net/http"
	"sale-system/src/controller"
	"sale-system/src/repository"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	repository.ConnectDB()

	registerProductRouter(router)

	http.ListenAndServe("localhost:8080", router)
}

func registerProductRouter(router *mux.Router) {
	router.HandleFunc("/products", controller.OptionsForBrowsers).Methods("OPTIONS")
	router.HandleFunc("/products", controller.FindAllProducts).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{code}", controller.FindProductById).Methods("GET")
}
