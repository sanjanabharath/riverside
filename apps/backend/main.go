package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanjanabharath/riverside/db"
	"github.com/sanjanabharath/riverside/models"
	"github.com/sanjanabharath/riverside/routes"
)

func main() {
	db.Connect()

	db.DB.AutoMigrate(&models.Task{})

	routes.RegisterRoutes()

	log.Println("Server running on server '8080'")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error while starting the server")
	}
}
