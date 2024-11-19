package main

import (
	"jobsprint/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Serve the index file (if needed)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "index.html")
	})

	// Serve adminpanel.html at /adminpanel route
	mux.HandleFunc("/adminpanel", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "static/adminpanel.html")
	})

	// Setup routes from handlers package
	handlers.SetupRoutes(mux)

	// CORS setup: allow all origins (you can restrict this to specific domains)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"}, // Replace "*" with specific domains to limit access
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Wrap your handler with the CORS handler
	handler := corsHandler.Handler(mux)

	// Start the server
	log.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", handler))

}
