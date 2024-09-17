package main

import (
	"api-server/internal/config"
	"api-server/internal/models"
	"api-server/internal/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{})

	router := routes.InitRoutes()

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS()(router)))
}
