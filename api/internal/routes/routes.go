package routes

import (
	"api/internal/handlers"
	"net/http"
)

func SetupRoutes(healthHandler *handlers.HealthHandler, albumHandler *handlers.AlbumHandler) {
	http.HandleFunc("/health", healthHandler.HealthCheck);

	http.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			albumHandler.GetAlbums(w, r);

		case http.MethodPost:
			albumHandler.CreateAlbum(w, r);

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
		}
	});
}
