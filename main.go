package main

import (
	"RenomachiBack/db"
	"RenomachiBack/handlers"
	"fmt"
	"net/http"
)

func main() {
	db.InitDB()

	http.HandleFunc("/users", handlers.HandleUsers)
	http.HandleFunc("/users/", handlers.HandleUser)
	http.HandleFunc("/buildings", handlers.HandleBuildings)
	http.HandleFunc("/buildings/", handlers.HandleBuilding)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
