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
func setupTestMusicianService(t *testing.T) services.MusicianServiceInterface {
	db := setupTestDB(t)
	repo := &repositories.MusicianRepository{DB: db}
	return &services.MusicianService{Repo: repo}
}

func TestCreateMusicianController(t *testing.T) {
	service := setupTestMusicianService(t)
	controller := &MusicianController{Service: service}

	musicianPayload := `{
		"name": "John Doe",
		"musician_type": "Guitarist"
	}`

	req := httptest.NewRequest("POST", "/musicians", bytes.NewBuffer([]byte(musicianPayload)))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	controller.CreateMusician(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %v, got %v", http.StatusCreated, rr.Code)
	}

	var musician models.Musician
	err := json.NewDecoder(rr.Body).Decode(&musician)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if musician.Name != "John Doe" {
		t.Errorf("expected musician name to be 'John Doe', got '%s'", musician.Name)
	}
}

func TestGetMusiciansController(t *testing.T) {
	service := setupTestMusicianService(t)
	controller := &MusicianController{Service: service}

	service.CreateMusician(&models.Musician{Name: "Musician 1", MusicianType: "Guitarist"})
	service.CreateMusician(&models.Musician{Name: "Musician 2", MusicianType: "Drummer"})

	req := httptest.NewRequest("GET", "/musicians", nil)
	rr := httptest.NewRecorder()

	controller.GetMusicians(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	var musicians []models.Musician
	err := json.NewDecoder(rr.Body).Decode(&musicians)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(musicians) != 2 {
		t.Errorf("expected 2 musicians, got %d", len(musicians))
	}
}

func TestUpdateMusicianController(t *testing.T) {
	service := setupTestMusicianService(t)
	controller := &MusicianController{Service: service}

	musician := &models.Musician{Name: "Old Musician", MusicianType: "Drummer"}
	service.CreateMusician(musician)

	updatedPayload := `{
		"name": "Updated Musician",
		"musician_type": "Guitarist"
	}`

	req := httptest.NewRequest("PUT", "/musicians/"+strconv.Itoa(int(musician.ID)), bytes.NewBuffer([]byte(updatedPayload)))
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(musician.ID))})
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	controller.UpdateMusician(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, rr.Code)
	}

	var updatedMusician models.Musician
	err := json.NewDecoder(rr.Body).Decode(&updatedMusician)
	if err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if updatedMusician.Name != "Updated Musician" {
		t.Errorf("expected musician name to be 'Updated Musician', got '%s'", updatedMusician.Name)
	}
}

func TestDeleteMusicianController(t *testing.T) {
	service := setupTestMusicianService(t)
	controller := &MusicianController{Service: service}

	musician := &models.Musician{Name: "Musician to Delete", MusicianType: "Bassist"}
	service.CreateMusician(musician)

	req := httptest.NewRequest("DELETE", "/musicians/"+strconv.Itoa(int(musician.ID)), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(int(musician.ID))})
	rr := httptest.NewRecorder()

	controller.DeleteMusician(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, got %v", http.StatusNoContent, rr.Code)
	}

	musicians, _ := service.GetMusicians()
	if len(musicians) != 0 {
		t.Errorf("expected 0 musicians, got %d", len(musicians))
	}
}
