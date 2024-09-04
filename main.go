package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// HelloHandler는 간단한 텍스트 응답을 반환합니다.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// GetUserHandler는 User 객체를 JSON으로 반환합니다.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Stella", Email: "stella@example.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// PostUserHandler는 POST 요청으로 받은 JSON 데이터를 처리합니다.
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "Received User: %s, Email: %s", user.Name, user.Email)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/user", GetUserHandler)    // GET 요청을 처리
	http.HandleFunc("/create", PostUserHandler) // POST 요청을 처리

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
