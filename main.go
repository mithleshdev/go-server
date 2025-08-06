package main

import (
	"log"
	"net/http"
	"time"
)

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		log.Printf("[%s] %s %s %s",
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the response time
		duration := time.Since(start)
		log.Printf("[%s] %s %s - Completed in %v",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}

func main() {
	mux := http.NewServeMux()

	// Health check endpoints for Kubernetes probes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Main endpoint
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Wrap the mux with logging middleware
	handler := loggingMiddleware(mux)

	log.Printf("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
