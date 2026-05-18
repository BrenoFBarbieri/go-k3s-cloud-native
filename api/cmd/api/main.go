package main

import (
	"api/internal/handlers"
	"api/internal/repositories"
	"api/internal/routes"
	"api/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	healthHandler := handlers.NewHealthHandler();

	albumRepository := repositories.NewAlbumRepository();
	albumService := services.NewAlbumService(albumRepository);
	albumHandler := handlers.NewAlbumHandler(albumService);

	routes.SetupRoutes(healthHandler, albumHandler);

	fmt.Println("Server running on port 8080");

	err := http.ListenAndServe(":8080", nil);
	if err != nil {
		log.Fatal(err);
	}
}
