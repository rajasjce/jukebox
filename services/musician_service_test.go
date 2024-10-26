package services_test

import (
	"database/sql"
	"jukebox/models"
	"jukebox/repositories"
	"jukebox/services"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestMusicianRepo(t *testing.T) *repositories.MusicianRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE musicians (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			musician_type TEXT NOT NULL
		);
		CREATE TABLE album_musicians (
			album_id INTEGER,
			musician_id INTEGER
		);
	`)
	if err != nil {
		t.Fatalf("failed to create test tables: %v", err)
	}

	return &repositories.MusicianRepository{DB: db}
}

func TestCreateMusicianService(t *testing.T) {
	repo := setupTestMusicianRepo(t)
	service := &services.MusicianService{Repo: repo}

	musician := &models.Musician{
		Name:         "Test Musician",
		MusicianType: "Guitarist",
	}

	err := service.CreateMusician(musician)
	if err != nil {
		t.Fatalf("failed to create musician: %v", err)
	}

	if musician.ID == 0 {
		t.Errorf("expected valid musician ID, got %d", musician.ID)
	}
}

func TestUpdateMusicianService(t *testing.T) {
	repo := setupTestMusicianRepo(t)
	service := &services.MusicianService{Repo: repo}

	// Insert a test musician
	musician := &models.Musician{
		Name:         "Old Musician",
		MusicianType: "Drummer",
	}
	err := service.CreateMusician(musician)
	if err != nil {
		t.Fatalf("failed to create musician: %v", err)
	}

	// Update the musician
	musician.Name = "Updated Musician"
	err = service.UpdateMusician(musician)
	if err != nil {
		t.Fatalf("failed to update musician: %v", err)
	}

	// Verify the update
	updatedMusician, err := repo.GetMusicians()
	if err != nil {
		t.Fatalf("failed to retrieve musicians: %v", err)
	}

	if updatedMusician[0].Name != "Updated Musician" {
		t.Errorf("expected 'Updated Musician', got '%s'", updatedMusician[0].Name)
	}
}

func TestDeleteMusicianService(t *testing.T) {
	repo := setupTestMusicianRepo(t)
	service := &services.MusicianService{Repo: repo}

	// Insert a test musician
	musician := &models.Musician{
		Name:         "Test Musician",
		MusicianType: "Pianist",
	}
	err := service.CreateMusician(musician)
	if err != nil {
		t.Fatalf("failed to create musician: %v", err)
	}

	// Delete the musician
	err = service.DeleteMusician(musician.ID)
	if err != nil {
		t.Fatalf("failed to delete musician: %v", err)
	}

	// Verify the deletion
	musicians, err := repo.GetMusicians()
	if err != nil {
		t.Fatalf("failed to get musicians: %v", err)
	}

	if len(musicians) != 0 {
		t.Errorf("expected 0 musicians, got %d", len(musicians))
	}
}

func TestGetMusiciansByAlbumService(t *testing.T) {
	repo := setupTestMusicianRepo(t)
	service := &services.MusicianService{Repo: repo}

	// Insert test musicians and link them to an album
	musician1 := &models.Musician{
		Name:         "Musician One",
		MusicianType: "Vocalist",
	}
	musician2 := &models.Musician{
		Name:         "Musician Two",
		MusicianType: "Bassist",
	}
	err := service.CreateMusician(musician1)
	if err != nil {
		t.Fatalf("failed to create musician 1: %v", err)
	}
	err = service.CreateMusician(musician2)
	if err != nil {
		t.Fatalf("failed to create musician 2: %v", err)
	}

	// Link musicians to album ID 1
	_, err = repo.DB.Exec(`INSERT INTO album_musicians (album_id, musician_id) VALUES (1, ?), (1, ?)`, musician1.ID, musician2.ID)
	if err != nil {
		t.Fatalf("failed to link musicians to album: %v", err)
	}

	// Get musicians by album
	musicians, err := service.GetMusiciansByAlbum(1)
	if err != nil {
		t.Fatalf("failed to get musicians by album: %v", err)
	}

	if len(musicians) != 2 {
		t.Errorf("expected 2 musicians, got %d", len(musicians))
	}
}

func TestGetMusiciansService(t *testing.T) {
	repo := setupTestMusicianRepo(t)
	service := &services.MusicianService{Repo: repo}

	// Insert test musicians
	repo.CreateMusician(&models.Musician{
		Name:         "Test Musician 1",
		MusicianType: "Guitarist",
	})
	repo.CreateMusician(&models.Musician{
		Name:         "Test Musician 2",
		MusicianType: "Drummer",
	})

	musicians, err := service.GetMusicians()
	if err != nil {
		t.Fatalf("failed to get musicians: %v", err)
	}

	if len(musicians) != 2 {
		t.Errorf("expected 2 musicians, got %d", len(musicians))
	}
}
