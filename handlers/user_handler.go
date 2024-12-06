package handlers

import (
	"RenomachiBack/utils"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	utils.PrintRequest(r)
	switch r.Method {
	case http.MethodGet:
		utils.TestResponseOK(w, r)
	case http.MethodPost:
		utils.TestResponseOK(w, r)
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
		case http.MethodGet:
			utils.TestResponseOK(w, r)
		case http.MethodDelete:
			utils.TestResponseOK(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	} else if len(params) > 2 && len(params) < 5 {
		if params[2] == "favorites" {
			if len(params) < 4 {
				switch r.Method {
				case http.MethodGet:
					utils.TestResponseOK(w, r)
				case http.MethodPost:
					utils.TestResponseOK(w, r)
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			} else if len(params) < 5 {
				switch r.Method {
				case http.MethodDelete:
					utils.TestResponseOK(w, r)
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}

			}
		} else if params[2] == "histories" {
			if len(params) < 4 {
				switch r.Method {
				case http.MethodGet:
					utils.TestResponseOK(w, r)
				default:
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			} else if len(params) < 5 {
				switch r.Method {
				case http.MethodDelete:
					utils.TestResponseOK(w, r)
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
