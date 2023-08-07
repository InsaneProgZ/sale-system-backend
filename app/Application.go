package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sale-system/controller"
	"sale-system/repository"
	"sale-system/service"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	controller, database, appConfig := setUp()
	defer database.Close()

	registerProductRouter(router, controller)

	http.ListenAndServe(fmt.Sprintf("%s:8080", appConfig.AppUrl), router)
}

func registerProductRouter(router *mux.Router, controller controller.Controller) {
	router.HandleFunc("/products", controller.OptionsForBrowsers).Methods("OPTIONS")
	router.HandleFunc("/products", controller.FindAllProducts).Methods("GET")
	router.HandleFunc("/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{code}", controller.FindProductByCode).Methods("GET")
	router.HandleFunc("/products/{code}", controller.ChangeProductByCode).Methods("PUT")
}

func setUp() (controller.Controller, *sql.DB, AppConfig) {
	var appConfig = getConfig()
	database := repository.ConnectDB(appConfig.DBUrl)
	repository := &repository.MysqlDB{Mysql: database}
	service := &service.ProductServiceImpl{Repository: repository}
	controller := &controller.ControllerImpl{Service: service}
	return controller, database, appConfig
}

func getConfig() AppConfig {
	var environment = os.Getenv("PROFILE")
	switch environment {
	case "container":
		{
			return AppConfig{AppUrl: "", DBUrl: "mysql"}
		}
	default:
		return AppConfig{AppUrl: "localhost", DBUrl: "localhost"}
	}
}

type AppConfig struct {
	AppUrl string
	DBUrl  string
}
