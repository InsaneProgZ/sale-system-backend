package web

import (
	"encoding/json"
	"net/http"
	"sale-system/src/adapter/input/web/request"
	"sale-system/src/adapter/input/web/router"
	"sale-system/src/adapter/output/repository"
	"sale-system/src/application/handler"
	"sale-system/src/application/service"
)

func Application() {
	http.HandleFunc("/products", router.CreateProductRouter)
	http.ListenAndServe("localhost:8081", nil)
}

type IProductController interface {
	handler(writer http.ResponseWriter, httpRequest *http.Request)
}

type productController struct {
	Service    service.IProductService
	Repository repository.IRespository
}

func newProductController(service service.IProductService, repository repository.IRespository) IProductController {
	return &productController{
		Service:    service,
		Repository: repository,
	}
}

func (pc *productController) handler(writer http.ResponseWriter, httpRequest *http.Request) {
	switch httpRequest.Method {
	case "POST":
		{
			var productRequest request.Product

			err := json.NewDecoder(httpRequest.Body).Decode(&productRequest)
			handler.ErrorHandler(err)

			product := productRequest.ToDomain()

			productResponse := pc.Service.CreateProduct(product).ToResponse()

			responseBody, err := json.Marshal(productResponse)
			handler.ErrorHandler(err)

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusCreated)
			writer.Write(responseBody)
		}
	default:
		http.Error(writer, "Not Found", http.StatusNotFound)

	}

}
