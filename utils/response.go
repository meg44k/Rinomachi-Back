package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONデータをレスポンス
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

func ResponseCreated(w http.ResponseWriter, data interface{}) {
	response := map[string]interface{}{
		"message": "User created successfully",
		"data":    data, // 作成されたユーザー情報を返す
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTPステータス201を設定
	json.NewEncoder(w).Encode(response)
}

func ResponseDeleted(w http.ResponseWriter, data string) {
	response := map[string]interface{}{
		"message": "User deleted successfully",
		"data":    data, // 削除されたユーザー情報を返す
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // HTTPステータス200を設定
	json.NewEncoder(w).Encode(response)
}

func ResponseUpdated(w http.ResponseWriter, data interface{}) {
	response := map[string]interface{}{
		"message": "User updated successfully",
		"data":    data, // 作成されたユーザー情報を返す
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTPステータス201を設定
	json.NewEncoder(w).Encode(response)
}
