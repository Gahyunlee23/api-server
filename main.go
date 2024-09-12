package main

import (
	"api-server/internal/routes"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	mux := routes.InitRoutes()

	handler := cors.Default().Handler(mux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
