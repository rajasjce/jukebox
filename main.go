package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"jukebox/controllers"
	"jukebox/repositories"
	"jukebox/routes"
	"jukebox/services"
)

func main() {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "./jukebox.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set up Repositories, Services, and Controllers
	albumRepo := &repositories.AlbumRepository{DB: db}
	musicianRepo := &repositories.MusicianRepository{DB: db}

	albumService := &services.AlbumService{Repo: albumRepo}
	musicianService := &services.MusicianService{Repo: musicianRepo}

	albumController := &controllers.AlbumController{Service: albumService}
	musicianController := &controllers.MusicianController{Service: musicianService}

	// Use mux for routing
	r := mux.NewRouter()

	// Set up the routes
	r = routes.SetupRoutes(albumController, musicianController)

	// Start the server
	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
