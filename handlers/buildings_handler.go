package handlers

import (
	"RenomachiBack/models"
	"RenomachiBack/utils"
	"net/http"
)

func HandleBuildings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		buildings, err := models.GetBuildings()
		if err != nil {
			http.Error(w, "Failed to fetch buildings", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, buildings, http.StatusOK)
	case http.MethodPost:
		utils.TestResponseOK(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleBuilding(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		utils.TestResponseOK(w, r)
	case http.MethodPost:
		utils.TestResponseOK(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
