package cmd

import (
	"log"
	"net/http"
	"nyctaxi_mapup/pkg/config"
	"nyctaxi_mapup/pkg/repository"
	"nyctaxi_mapup/pkg/router"
)

func main() {
	router := router.SetupRoutes()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func init() {
	config.LoadConfig()

	err := repository.InitializeDB()
	if err != nil {
		log.Panic("error while connecting to database: ", err)
	}
}
