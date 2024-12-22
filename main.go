package main

import (
	"RenomachiBack/db"
	"RenomachiBack/handlers"
	"RenomachiBack/utils"
	"fmt"
	"net/http"
)

func main() {
	db.InitDB()

	http.Handle("/users", utils.EnableCORS(http.HandlerFunc(handlers.HandleUsers)))
	http.Handle("/users/", utils.EnableCORS(http.HandlerFunc(handlers.HandleUser)))
	http.Handle("/buildings", utils.EnableCORS(http.HandlerFunc(handlers.HandleBuildings)))
	http.Handle("/buildings/", utils.EnableCORS(http.HandlerFunc(handlers.HandleBuilding)))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
