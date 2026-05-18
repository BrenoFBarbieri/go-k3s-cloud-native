package repositories

import "api/internal/models"

type AlbumRepository struct{}

var albums = []models.Album{
	{
		ID: "1",
		Title: "Blue Train",
		Artist: "John Coltrane",
		Price: 56.99,
	},
	{
		ID: "2",
		Title: "Jeru",
		Artist: "Gerry Mulligan",
		Price: 17.99,
	},
}

func NewAlbumRepository() *AlbumRepository {
	return &AlbumRepository{};
}

func (r *AlbumRepository) FindAll() []models.Album {
	return albums;
}

func (r *AlbumRepository) Create(album models.Album) {
	albums = append(albums, album);
}
