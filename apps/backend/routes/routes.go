package routes

import (
	"net/http"

	"github.com/sanjanabharath/riverside/handlers"
)


func RegisterRoutes() {
	http.HandleFunc("POST	/createTask", handlers.CreateTask)
}