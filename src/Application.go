package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/InsaneProgZ/sale-system-backend/adapter/input/controller"
	"github.com/InsaneProgZ/sale-system-backend/adapter/output/mysql"
	"github.com/InsaneProgZ/sale-system-backend/domain/service"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	controller, database, appConfig := setUp()
	defer database.Close()

	registerProductRouter(router, controller)
	http.ListenAndServe(fmt.Sprintf("%s:8080", appConfig.AppUrl), router)
}

const productResource = "/products"

func registerProductRouter(router *mux.Router, controller controller.Controller) {
	router.HandleFunc(productResource, controller.OptionsForBrowsers).Methods("OPTIONS")
	router.HandleFunc(productResource, controller.FindAllProducts).Methods("GET")
	router.HandleFunc(productResource, controller.CreateProduct).Methods("POST")
	router.HandleFunc(productResource + "/{code}", controller.FindProductByCode).Methods("GET")
	router.HandleFunc(productResource + "/{code}", controller.ChangeProductByCode).Methods("PUT")
}

func setUp() (controller.Controller, *sql.DB, AppConfig) {
	var appConfig = getConfig()
	database := mysql.ConnectDB(appConfig.DBUrl)
	mysql := &mysql.MysqlDB{Mysql: database}
	service := &service.ProductServiceImpl{Repository: mysql}
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
