package controllers

import (
	"bytes"
	"encoding/json"
	"jukebox/models"
	"jukebox/repositories"
	"jukebox/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Helper function to setup an in-memory service for testing purposes
func setupTestService(t *testing.T) services.AlbumServiceInterface {
	db := setupTestDB(t)

	repo := &repositories.AlbumRepository{DB: db}
	return &services.AlbumService{Repo: repo}
}

func TestCreateAlbumController(t *testing.T) {
	service := setupTestService(t)
	controller := &AlbumController{Service: service}

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

	controller.CreateAlbum(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %v, got %v", http.StatusCreated, rr.Code)
	}

	var album models.Album
	err := json.NewDecoder(rr.Body).Decode(&album)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if album.Name != "Test Album" {
		t.Errorf("expected album name to be 'Test Album', got '%s'", album.Name)
	}
}

func TestGetAlbumsController(t *testing.T) {
	service := setupTestService(t)
	controller := &AlbumController{Service: service}

	service.CreateAlbum(&models.Album{Name: "Album 1", ReleaseDate: "2022-01-01", Genre: "Rock", Price: 150, Description: "Description 1"})
	service.CreateAlbum(&models.Album{Name: "Album 2", ReleaseDate: "2022-02-01", Genre: "Pop", Price: 200, Description: "Description 2"})

	req := httptest.NewRequest("GET", "/albums", nil)
	rr := httptest.NewRecorder()

	controller.GetAlbums(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	var albums []models.Album
	err := json.NewDecoder(rr.Body).Decode(&albums)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(albums) != 2 {
		t.Errorf("expected 2 albums, got %d", len(albums))
	}
}

func TestUpdateAlbumController(t *testing.T) {
	service := setupTestService(t)
	controller := &AlbumController{Service: service}

	album := &models.Album{Name: "Old Album", ReleaseDate: "2022-01-01", Genre: "Rock", Price: 150, Description: "Old Description"}
	service.CreateAlbum(album)

	updatedPayload := `{
		"name": "Updated Album",
		"release_date": "2022-01-01",
		"genre": "Rock",
		"price": 180,
		"description": "Updated Description"
	}`

	req := httptest.NewRequest("PUT", "/albums/"+strconv.Itoa(int(album.ID)), bytes.NewBuffer([]byte(updatedPayload)))
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(album.ID))})
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	controller.UpdateAlbum(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	var updatedAlbum models.Album
	err := json.NewDecoder(rr.Body).Decode(&updatedAlbum)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if updatedAlbum.Name != "Updated Album" {
		t.Errorf("expected album name to be 'Updated Album', got '%s'", updatedAlbum.Name)
	}
}

func TestDeleteAlbumController(t *testing.T) {
	service := setupTestService(t)
	controller := &AlbumController{Service: service}

	album := &models.Album{Name: "Album to Delete", ReleaseDate: "2022-01-01", Genre: "Rock", Price: 150, Description: "Description"}
	service.CreateAlbum(album)

	req := httptest.NewRequest("DELETE", "/albums/"+strconv.Itoa(int(album.ID)), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(album.ID))})
	rr := httptest.NewRecorder()

	controller.DeleteAlbum(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, got %v", http.StatusNoContent, rr.Code)
	}

	albums, _ := service.GetAlbums()
	if len(albums) != 0 {
		t.Errorf("expected 0 albums, got %d", len(albums))
	}
}

func TestGetAlbumsByMusicianController(t *testing.T) {
	service := setupTestService(t)
	controller := &AlbumController{Service: service}

	album1 := &models.Album{Name: "Rock Album", ReleaseDate: "2022-01-01", Genre: "Rock", Price: 200, Description: "First Album"}
	album2 := &models.Album{Name: "Pop Album", ReleaseDate: "2022-02-01", Genre: "Pop", Price: 300, Description: "Second Album"}

	service.CreateAlbum(album1)
	service.CreateAlbum(album2)
	service.LinkMusiciansToAlbum(album1.ID, []uint{101})
	service.LinkMusiciansToAlbum(album2.ID, []uint{101})

	req := httptest.NewRequest("GET", "/musicians/101/albums", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "101"})
	rr := httptest.NewRecorder()

	controller.GetAlbumsByMusician(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	var albums []models.Album
	err := json.NewDecoder(rr.Body).Decode(&albums)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(albums) != 2 {
		t.Errorf("expected 2 albums for musician 101, got %d", len(albums))
	}
}
