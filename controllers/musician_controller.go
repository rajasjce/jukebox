package controllers

import (
	"encoding/json"
	"jukebox/models"
	"jukebox/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MusicianController struct {
	Service services.MusicianServiceInterface
}

// CreateMusician handles the creation of a new musician.
func (c *MusicianController) CreateMusician(w http.ResponseWriter, r *http.Request) {
	var musician models.Musician
	if err := json.NewDecoder(r.Body).Decode(&musician); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateMusician(&musician); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(musician)
}

// GetMusicians handles retrieving a list of musicians.
func (c *MusicianController) GetMusicians(w http.ResponseWriter, r *http.Request) {
	musicians, err := c.Service.GetMusicians()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(musicians)
}

// UpdateMusician handles updating an existing musician.
func (c *MusicianController) UpdateMusician(w http.ResponseWriter, r *http.Request) {
	var musician models.Musician
	if err := json.NewDecoder(r.Body).Decode(&musician); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse musician ID from the URL
	vars := mux.Vars(r)
	musicianID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid musician ID", http.StatusBadRequest)
		return
	}

	musician.ID = uint(musicianID)

	if err := c.Service.UpdateMusician(&musician); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(musician)
}

// DeleteMusician handles deleting a musician by ID.
func (c *MusicianController) DeleteMusician(w http.ResponseWriter, r *http.Request) {
	// Parse musician ID from the URL
	vars := mux.Vars(r)
	musicianID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid musician ID", http.StatusBadRequest)
		return
	}

	if err := c.Service.DeleteMusician(uint(musicianID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetMusiciansByAlbum handles retrieving musicians for a specific album.
func (c *MusicianController) GetMusiciansByAlbum(w http.ResponseWriter, r *http.Request) {
	// Parse album ID from the URL
	vars := mux.Vars(r)
	albumID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid album ID", http.StatusBadRequest)
		return
	}

	musicians, err := c.Service.GetMusiciansByAlbum(uint(albumID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(musicians)
}
