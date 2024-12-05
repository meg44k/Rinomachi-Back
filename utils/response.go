package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse sends a JSON response
func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func TestResponseOK(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("%s requested!\n", r.URL.String())
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
