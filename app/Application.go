package main

import (
	"database/sql"
	"net/http"
	"sale-system/controller"
	"sale-system/repository"
	"sale-system/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	controller, database := setUp()
	defer database.Close()

	registerProductRouter(router, controller)

	http.ListenAndServe(":8080", router)
}

func registerProductRouter(router *mux.Router, controller controller.Controller) {
	router.HandleFunc("/products", controller.OptionsForBrowsers).Methods("OPTIONS")
	router.HandleFunc("/products", controller.FindAllProducts).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{code}", controller.FindProductByCode).Methods("GET")
	router.HandleFunc("/products/{code}", controller.ChangeProductByCode).Methods("PUT")
}

func setUp() (controller.Controller, *sql.DB) {
	database := repository.ConnectDB()
	repository := &repository.MysqlDB{Mysql: database}
	service := &service.ProductServiceImpl{Repository: repository}
	controller := &controller.ControllerImpl{Service: service}
	return controller, database
}
