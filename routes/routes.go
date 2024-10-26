package routes

import (
	"jukebox/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(albumController *controllers.AlbumController, musicianController *controllers.MusicianController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Jukebox API"))
	})

	// Define Routes for Albums and Musicians
	r.HandleFunc("/albums", albumController.GetAlbums).Methods("GET")
	r.HandleFunc("/albums", albumController.CreateAlbum).Methods("POST")
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.UpdateAlbum).Methods("PUT")                   // Update album by ID
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.DeleteAlbum).Methods("DELETE")                // Delete album by ID
	r.HandleFunc("/musicians/{id:[0-9]+}/albums", albumController.GetAlbumsByMusician).Methods("GET") // Get albums by musician ID

	r.HandleFunc("/musicians", musicianController.GetMusicians).Methods("GET")
	r.HandleFunc("/musicians", musicianController.CreateMusician).Methods("POST")
	r.HandleFunc("/musicians/{id:[0-9]+}", musicianController.UpdateMusician).Methods("PUT")             // Update musician by ID
	r.HandleFunc("/musicians/{id:[0-9]+}", musicianController.DeleteMusician).Methods("DELETE")          // Delete musician by ID
	r.HandleFunc("/albums/{id:[0-9]+}/musicians", musicianController.GetMusiciansByAlbum).Methods("GET") // Get musicians by album ID

	return r
}
