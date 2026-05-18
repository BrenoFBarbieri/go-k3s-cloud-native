package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "ok",
		"service": "album-api",
		"version": "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusOK);

	json.NewEncoder(w).Encode(response);
}
