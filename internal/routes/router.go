package routes

import (
	"api-server/internal/controllers"
	_ "net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/products", controllers.CreateProductHandler).Methods("POST")
	router.HandleFunc("/products", controllers.GetAllProductsHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.GetProductHandler).Methods("GET")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.UpdateProductHandler).Methods("PUT")
	router.HandleFunc("/products/{id:[0-9]+}", controllers.DeleteProductHandler).Methods("DELETE")

	return router
}
