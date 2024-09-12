package routes

import (
	"api-server/internal/controllers"
	"net/http"
)

func InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("/user", controllers.GetUserHandler)
	mux.HandleFunc("/create", controllers.PostUserHandler)
	mux.HandleFunc("/product", controllers.GetProductHandler)

	return mux
}
