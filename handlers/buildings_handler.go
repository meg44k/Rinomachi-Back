package handlers

import (
	"RenomachiBack/utils"
	"net/http"
)

func HandleBuildings(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		utils.TestResponseOK(w, r)
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
