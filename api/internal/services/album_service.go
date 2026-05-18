package services

import (
	"api/internal/models"
	"api/internal/repositories"
);

type AlbumService struct {
	repository *repositories.AlbumRepository;
}

func NewAlbumService(r *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{ repository: r };
}

func (s *AlbumService) GetAlbums() []models.Album {
	return s.repository.FindAll();
}

func (s *AlbumService) CreateAlbum(album models.Album) {
	s.repository.Create(album);
}
