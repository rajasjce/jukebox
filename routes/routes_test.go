package routes

import (
	"bytes"
	"encoding/json"
	"jukebox/controllers"
	"jukebox/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesSetup(t *testing.T) {
	// Use the in-memory services
	albumService := &InMemoryAlbumService{}
	musicianService := &InMemoryMusicianService{}

	// Create controllers with the in-memory services
	albumController := &controllers.AlbumController{Service: albumService}
	musicianController := &controllers.MusicianController{Service: musicianService}

	// Setup routes
	router := SetupRoutes(albumController, musicianController)

	// Test CreateAlbum Route
	albumPayload := `{
		"name": "Test Album",
		"release_date": "2022-01-01",
		"genre": "Rock",
		"price": 150,
		"description": "A test album",
		"musician_ids": [1, 2]
	}`

	req := httptest.NewRequest("POST", "/albums", bytes.NewBuffer([]byte(albumPayload)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %v, got %v", http.StatusCreated, rr.Code)
	}

	// Parse the response body
	var response models.Album
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Errorf("failed to parse response: %v", err)
	}

	// Check if the album has an ID
	if response.ID == 0 {
		t.Errorf("expected album to have an ID, got %v", response.ID)
	}

	// Test GetAlbums Route
	req = httptest.NewRequest("GET", "/albums", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}
}
