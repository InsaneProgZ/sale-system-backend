package main

import (
	"database/sql"
	"net/http"
	"sale-system/src/controller"
	"sale-system/src/repository"
	"sale-system/src/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	controller, database := setUp()
	defer database.Close()

	registerProductRouter(router, controller)

	http.ListenAndServe("localhost:8080", router)
}

func registerProductRouter(router *mux.Router, controller controller.Controller) {
	router.HandleFunc("/products", controller.OptionsForBrowsers).Methods("OPTIONS")
	router.HandleFunc("/products", controller.FindAllProducts).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{code}", controller.FindProductById).Methods("GET")
}

func setUp() (controller.Controller, *sql.DB) {
	database := repository.ConnectDB()
	repository := &repository.MysqlDB{Mysql: database}
	service := &service.ProductServiceImpl{Repository: repository}
	controller := &controller.ControllerImpl{Service: service}
	return controller, database
}
