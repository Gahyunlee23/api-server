package controllers

import (
	"api-server/internal/models"
	"encoding/json"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{Id: 70, Name: "\tStandard Business card", Code: "pp_business_card", Type: "upload_fullpage_configuration"},
		{Id: 40, Name: "\tLabels (Cut Sheet)", Code: "pp_labelscutsheet_np", Type: "upload_personalization_fullpage_configuration"},
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}
