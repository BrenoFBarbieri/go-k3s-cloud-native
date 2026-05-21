package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "ok",
		"service":   os.Getenv("APP_NAME"),
		"version":   os.Getenv("APP_VERSION"),
		"env":       os.Getenv("APP_ENV"),
		"jwtSecret": os.Getenv("JWT_SECRET"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Simulate an error for testing

	json.NewEncoder(w).Encode(response)
}
