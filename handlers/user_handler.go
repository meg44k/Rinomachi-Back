package handlers

import (
	"RenomachiBack/models"
	"RenomachiBack/utils"
	"encoding/json"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	utils.PrintRequest(r)
	switch r.Method {
	// GETメソッド
	case http.MethodGet:
		users, err := models.GetUsers()
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		utils.JSONResponse(w, users, http.StatusOK)
	// POSTメソッド
	case http.MethodPost:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Failed to insert user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = user.AddUser()
		if err != nil {
			http.Error(w, "Failed to insert user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		utils.ResponseCreated(w, user)
	// 許可されていないメソッド
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	// リクエストのMethodとURLを表示
	utils.PrintRequest(r)
	// ルートパラメータの取得 (例: /users/1/favorites -> ["users","1","favorites"])
	params := utils.GetRouteParams(r)

	// /users/のあとにパラメータが指定されていない時
	if len(params) < 2 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if len(params) == 2 {
		switch r.Method {
		// GETメソッド
		case http.MethodGet:
			user, err := models.GetUser(params[1])
			if err != nil {
				http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
				return
			}
			utils.JSONResponse(w, user, http.StatusOK)
		// DELETEメソッド
		case http.MethodDelete:
			err := models.DeleteUser(params[1])
			if err != nil {
				http.Error(w, "Failed to delete user", http.StatusInternalServerError)
				return
			}
			utils.ResponseDeleted(w, params[1])
		// PUTメソッド
		case http.MethodPut:
			var user models.User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
				return
			}
			err = models.UpdateUser(&user, params[1])
			if err != nil {
				http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
				return
			}
			utils.ResponseUpdated(w, user)
		// 許可されていないメソッド
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	} else if len(params) > 2 && len(params) < 5 {
		if params[2] == "favorites" {
			if len(params) < 4 {
				switch r.Method {
				// GETメソッド
				case http.MethodGet:
					favorites, err := models.GetFavorites(params[1])
					if err != nil {
						http.Error(w, "Failed to fetch favorites", http.StatusInternalServerError)
						return
					}
					utils.JSONResponse(w, favorites, http.StatusOK)
				// POSTメソッド
				case http.MethodPost:
					var favorite models.Favorite
					err := json.NewDecoder(r.Body).Decode(&favorite)
					if err != nil {
						http.Error(w, "Failed to insert favorite: "+err.Error(), http.StatusInternalServerError)
						return
					}
					err = favorite.AddFavorite()
					if err != nil {
						http.Error(w, "Failed to insert favorite: "+err.Error(), http.StatusInternalServerError)
						return
					}
					utils.ResponseCreated(w, favorite)
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			} else if len(params) < 5 {
				switch r.Method {
				// DELETEメソッド
				case http.MethodDelete:
					err := models.DeleteFavorite(params[1], params[3])
					if err != nil {
						http.Error(w, "Failed to delete favorites", http.StatusInternalServerError)
						return
					}
					utils.ResponseDeleted(w, params[3])
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}

			}
		} else if params[2] == "histories" {
			if len(params) < 4 {
				switch r.Method {
				// GETメソッド
				case http.MethodGet:
					histories, err := models.GetHistories(params[1])
					if err != nil {
						http.Error(w, "Failed to fetch histories", http.StatusInternalServerError)
						return
					}
					utils.JSONResponse(w, histories, http.StatusOK)
				// POSTメソッド
				case http.MethodPost:
					var history models.History
					err := json.NewDecoder(r.Body).Decode(&history)
					if err != nil {
						http.Error(w, "Failed to insert history: "+err.Error(), http.StatusInternalServerError)
						return
					}
					err = history.AddHistory()
					if err != nil {
						http.Error(w, "Failed to insert history: "+err.Error(), http.StatusInternalServerError)
						return
					}
					utils.ResponseCreated(w, history)
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			} else if len(params) < 5 {
				switch r.Method {
				case http.MethodDelete:
					err := models.DeleteHistory(params[1], params[3])
					if err != nil {
						http.Error(w, "Failed to delete history: "+err.Error(), http.StatusInternalServerError)
						return
					}
					utils.ResponseDeleted(w, params[3])
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			}
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
