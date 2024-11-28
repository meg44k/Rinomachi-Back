package handlers

import "net/http"

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
