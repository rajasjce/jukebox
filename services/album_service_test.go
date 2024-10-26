package services_test

import (
	"database/sql"
	"jukebox/models"
	"jukebox/repositories"
	"jukebox/services"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestRepo(t *testing.T) *repositories.AlbumRepository {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE albums (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			release_date TEXT,
			genre TEXT,
			price REAL,
			description TEXT
		);
		CREATE TABLE album_musicians (
			album_id INTEGER,
			musician_id INTEGER
		);
	`)
	if err != nil {
		t.Fatalf("failed to create test tables: %v", err)
	}

	return &repositories.AlbumRepository{DB: db}
}

func TestCreateAlbumService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	album := &models.Album{
		Name:        "Test Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "A test album",
	}

	err := service.CreateAlbum(album)
	if err != nil {
		t.Fatalf("failed to create album: %v", err)
	}

	if album.ID == 0 {
		t.Errorf("expected valid album ID, got %d", album.ID)
	}
}

func TestUpdateAlbumService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	// Insert a test album
	album := &models.Album{
		Name:        "Old Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "An old album",
	}
	err := service.CreateAlbum(album)
	if err != nil {
		t.Fatalf("failed to create album: %v", err)
	}

	// Update the album
	album.Name = "Updated Album"
	album.Price = 220
	err = service.UpdateAlbum(album)
	if err != nil {
		t.Fatalf("failed to update album: %v", err)
	}

	// Verify the update
	updatedAlbum, err := repo.GetAlbums()
	if err != nil {
		t.Fatalf("failed to retrieve albums: %v", err)
	}

	if updatedAlbum[0].Name != "Updated Album" {
		t.Errorf("expected 'Updated Album', got '%s'", updatedAlbum[0].Name)
	}
}

func TestDeleteAlbumService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	// Insert a test album
	album := &models.Album{
		Name:        "Test Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "A test album",
	}
	err := service.CreateAlbum(album)
	if err != nil {
		t.Fatalf("failed to create album: %v", err)
	}

	// Delete the album
	err = service.DeleteAlbum(album.ID)
	if err != nil {
		t.Fatalf("failed to delete album: %v", err)
	}

	// Verify the deletion
	albums, err := repo.GetAlbums()
	if err != nil {
		t.Fatalf("failed to get albums: %v", err)
	}

	if len(albums) != 0 {
		t.Errorf("expected 0 albums, got %d", len(albums))
	}
}

func TestLinkMusiciansToAlbumService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	// Insert a test album
	album := &models.Album{
		Name:        "Test Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "A test album",
	}
	err := service.CreateAlbum(album)
	if err != nil {
		t.Fatalf("failed to create album: %v", err)
	}

	// Link musicians to the album
	err = service.LinkMusiciansToAlbum(album.ID, []uint{101, 102})
	if err != nil {
		t.Fatalf("failed to link musicians to album: %v", err)
	}

	// Verify the link
	var count int
	db := repo.DB
	err = db.QueryRow("SELECT COUNT(*) FROM album_musicians WHERE album_id = ?", album.ID).Scan(&count)
	if err != nil {
		t.Fatalf("failed to count album_musicians: %v", err)
	}

	if count != 2 {
		t.Errorf("expected 2 musicians linked to album, got %d", count)
	}
}

func TestGetAlbumsByMusicianService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	// Insert test albums and musicians
	album1 := &models.Album{
		Name:        "Rock Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "First Album",
	}
	album2 := &models.Album{
		Name:        "Pop Album",
		ReleaseDate: "2022-02-01",
		Genre:       "Pop",
		Price:       300,
		Description: "Second Album",
	}

	err := service.CreateAlbum(album1)
	if err != nil {
		t.Fatalf("failed to create album 1: %v", err)
	}
	err = service.CreateAlbum(album2)
	if err != nil {
		t.Fatalf("failed to create album 2: %v", err)
	}

	err = service.LinkMusiciansToAlbum(album1.ID, []uint{101})
	if err != nil {
		t.Fatalf("failed to link musicians to album 1: %v", err)
	}
	err = service.LinkMusiciansToAlbum(album2.ID, []uint{101})
	if err != nil {
		t.Fatalf("failed to link musicians to album 2: %v", err)
	}

	// Test GetAlbumsByMusician
	albums, err := service.GetAlbumsByMusician(101)
	if err != nil {
		t.Fatalf("failed to get albums by musician: %v", err)
	}

	// Verify the number of albums retrieved
	if len(albums) != 2 {
		t.Errorf("expected 2 albums for musician 101, got %d", len(albums))
	}
}

func TestGetAlbumsService(t *testing.T) {
	repo := setupTestRepo(t)
	service := &services.AlbumService{Repo: repo}

	// Insert test data
	repo.CreateAlbum(&models.Album{
		Name:        "Test Album 1",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "A test album 1",
	})
	repo.CreateAlbum(&models.Album{
		Name:        "Test Album 2",
		ReleaseDate: "2022-02-01",
		Genre:       "Pop",
		Price:       250,
		Description: "A test album 2",
	})

	albums, err := service.GetAlbums()
	if err != nil {
		t.Fatalf("failed to get albums: %v", err)
	}

	if len(albums) != 2 {
		t.Errorf("expected 2 albums, got %d", len(albums))
	}
}
