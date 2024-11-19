package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// SetupRoutes sets up HTTP routes and handlers
func SetupRoutes(mux *http.ServeMux) {
	apiMux := http.NewServeMux()

	// Define Routes
	apiMux.HandleFunc("/submit", handleFormSubmission)
	apiMux.HandleFunc("/users", handleGetAllUsers)

	mux.Handle("/api/", restrictAPIAccess(http.StripPrefix("/api", apiMux)))
}

func restrictAPIAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		referer := r.Header.Get("Referer")

		if strings.Contains(referer, "/api/") || referer == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"error": "Access from " + referer + " is restricted"})
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
