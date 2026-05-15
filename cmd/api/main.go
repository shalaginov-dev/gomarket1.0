package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Health struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{Status: "ok"}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	fmt.Println("Server started on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		fmt.Println("Server failed:", err)
	}
}
