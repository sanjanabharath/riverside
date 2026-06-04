package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sanjanabharath/riverside/db"
	"github.com/sanjanabharath/riverside/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
	var task []models.Task

	result := db.DB.Find(&task) // SELECT * from tasks

	if result.Error != nil {
		http.Error(w, "Failed to get any task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}