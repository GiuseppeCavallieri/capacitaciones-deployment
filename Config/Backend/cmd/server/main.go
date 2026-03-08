package main

import (
	"fmt"

	"net/http"
	"veterinaria/backend/internal/config"
	"veterinaria/backend/internal/database"
)

func main() {
	config := config.LoadConfig()
	session, err := database.ConnectDB(config)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer session.Close()

	http.ListenAndServe(config.Port, nil)

	fmt.Println("Connected to database successfully")
}