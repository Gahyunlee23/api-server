package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Id: 70, Name: "\tStandard Business card", Code: "pp_business_card"},
		{Id: 40, Name: "\tLabels (Cut Sheet)", Code: "pp_labelscutsheet_np"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Stella", Email: "stella@example.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "Received User: %s, Email: %s", user.Name, user.Email)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/user", GetUserHandler)
	mux.HandleFunc("/create", PostUserHandler)
	mux.HandleFunc("/product", GetProductHandler)

	// CORS 허용 설정
	handler := cors.Default().Handler(mux)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
