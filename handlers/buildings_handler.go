package handlers

import (
	"RenomachiBack/models"
	"RenomachiBack/utils"
	"encoding/json"
	"net/http"
)

func HandleBuildings(w http.ResponseWriter, r *http.Request) {
	utils.PrintRequest(r)
	switch r.Method {
	case http.MethodGet:
		buildings, err := models.GetBuildings()
		if err != nil {
			http.Error(w, "Failed to fetch buildings", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, buildings, http.StatusOK)
	case http.MethodPost:
		var building models.Building
		err := json.NewDecoder(r.Body).Decode(&building)
		if err != nil {
			http.Error(w, "Failed to insert building: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = building.AddBuilding()
		if err != nil {
			http.Error(w, "Failed to insert building: "+err.Error(), http.StatusInternalServerError)
			return
		}
		utils.ResponseCreated(w, building)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleBuilding(w http.ResponseWriter, r *http.Request) {
	utils.PrintRequest(r)
	params := utils.GetRouteParams(r)
	if len(params) < 2 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if len(params) == 2 {
		switch r.Method {
		case http.MethodGet:
			utils.TestResponseOK(w, r)
		case http.MethodPost:
			utils.TestResponseOK(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
