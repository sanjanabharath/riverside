package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sanjanabharath/riverside/db"
	"github.com/sanjanabharath/riverside/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	result := db.DB.Create(&task)

	if result.Error != nil {
		http.Error(w, "Failed to create a task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}