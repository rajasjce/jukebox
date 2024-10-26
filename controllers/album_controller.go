package controllers

import (
	"encoding/json"
	"jukebox/models"
	"jukebox/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AlbumController struct {
	Service services.AlbumServiceInterface
}

var albumDTO struct {
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	Genre       string  `json:"genre"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	MusicianIDs []uint  `json:"musician_ids"`
}

func (c *AlbumController) CreateAlbum(w http.ResponseWriter, r *http.Request) {

	// Decode the request body into albumDTO
	if err := json.NewDecoder(r.Body).Decode(&albumDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	album := models.Album{
		Name:        albumDTO.Name,
		ReleaseDate: albumDTO.ReleaseDate,
		Genre:       albumDTO.Genre,
		Price:       albumDTO.Price,
		Description: albumDTO.Description,
	}

	// Insert album into the albums table
	if err := c.Service.CreateAlbum(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Automatically link album to musicians by calling a helper function
	if err := c.Service.LinkMusiciansToAlbum(album.ID, albumDTO.MusicianIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(album)
}

// GetAlbums handles retrieving all albums from the database.
func (c *AlbumController) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := c.Service.GetAlbums()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}

// UpdateAlbum handles updating an existing album.
func (c *AlbumController) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse album ID from the request URL
	vars := mux.Vars(r)
	albumID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	album.ID = uint(albumID)

	if err := c.Service.UpdateAlbum(&album); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(album)
}

// DeleteAlbum handles deleting an album by ID.
func (c *AlbumController) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	// Parse album ID from the request URL
	vars := mux.Vars(r)
	albumID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	if err := c.Service.DeleteAlbum(uint(albumID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAlbumsByMusician handles retrieving albums for a specific musician, sorted by price.
func (c *AlbumController) GetAlbumsByMusician(w http.ResponseWriter, r *http.Request) {
	// Parse musician ID from the request URL
	vars := mux.Vars(r)
	musicianID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid musician ID", http.StatusBadRequest)
		return
	}

	albums, err := c.Service.GetAlbumsByMusician(uint(musicianID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(albums)
}
