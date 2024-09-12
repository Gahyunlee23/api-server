package controllers

import (
	"api-server/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: "Stella", Email: "stella@example.com"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Fprintf(w, "Received User: %s, Email: %s", user.Name, user.Email)
}
