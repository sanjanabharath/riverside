package routes

import (
	"net/http"

	"github.com/sanjanabharath/riverside/handlers"
	"github.com/sanjanabharath/riverside/utils"
)


func RegisterRoutes() {
	http.HandleFunc("/createTask", utils.EnableCORS(handlers.CreateTask))
	http.HandleFunc("/getTasks", utils.EnableCORS(handlers.GetUsers))
}