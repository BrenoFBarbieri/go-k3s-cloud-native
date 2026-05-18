package handlers

import (
	"api/internal/models"
	"api/internal/services"
	"encoding/json"
	"net/http"
)

type AlbumHandler struct {
	service *services.AlbumService
}

func NewAlbumHandler(service *services.AlbumService) *AlbumHandler {
	return &AlbumHandler{ service: service }
}


func (h *AlbumHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json");

	albums := h.service.GetAlbums();

	json.NewEncoder(w).Encode(albums);
}

func (h *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album;

	err := json.NewDecoder(r.Body).Decode(&album);
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	h.service.CreateAlbum(album);

	w.WriteHeader(http.StatusCreated);

	json.NewEncoder(w).Encode(album);
}
