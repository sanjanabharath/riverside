package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", enableCORS(homeHandler))

	log.Println("Server running on server '8080'")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error while starting the server")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func addTask(w http.ResponseWriter, r *http.Request) {
	
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		//preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}