package controllers

import (
	"api-server/internal/models"
	"encoding/json"
	"net/http"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{Id: 70, Name: "\tStandard Business card", Code: "BRCHRSMLNG1", Type: "upload_fullpage_configuration"},
		{Id: 40, Name: "\tLabels (Cut Sheet)", Code: "PSTRCDSMLNG4", Type: "\tupload_fullpage_configuration"},
		{Id: 26, Name: "\tPostcards with Direct Mail by Canada Post", Code: "PSTRCDSMLNG3", Type: "\tupload_fullpage_configuration"},
		{Id: 27, Name: "\t4D Lenticular Cards", Code: "4DLCARDS5", Type: "upload_personalization_fullpage_configuration"},
		{Id: 28, Name: "\t\tBookmarks", Code: "BKMRKS6", Type: "\tupload_fullpage_configuration"},
		{Id: 29, Name: "\tBottle Neck Hang Tags", Code: "BTTLNECKHT7", Type: "\tupload_fullpage_configuration"},
		{Id: 30, Name: "\tBrochures", Code: "BRCHR8", Type: "\tupload_fullpage_configuration"},
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		return
	}
}
