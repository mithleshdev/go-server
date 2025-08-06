package main

import (
	"log"
	"net/http"
)

func main() {
	// Health check endpoints for Kubernetes probes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Main endpoint
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
