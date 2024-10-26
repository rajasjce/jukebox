package repositories

import (
	"jukebox/models"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateAlbum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	album := &models.Album{
		Name:        "Test Album",
		ReleaseDate: "2022-01-01",
		Genre:       "Rock",
		Price:       200,
		Description: "A test album",
	}

	err := repo.CreateAlbum(album)
	if err != nil {
		t.Fatalf("failed to create album: %v", err)
	}

	if album.ID == 0 {
		t.Errorf("expected valid album ID, got %d", album.ID)
	}
}

func TestGetAlbums(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	// Insert test data
	_, err := db.Exec(`
		INSERT INTO albums (name, release_date, genre, price, description) VALUES 
		('Test Album 1', '2022-01-01', 'Rock', 200, 'A test album 1'),
		('Test Album 2', '2022-02-01', 'Pop', 250, 'A test album 2');
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	albums, err := repo.GetAlbums()
	if err != nil {
		t.Fatalf("failed to get albums: %v", err)
	}

	if len(albums) != 2 {
		t.Errorf("expected 2 albums, got %d", len(albums))
	}
}

func TestUpdateAlbum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	// Insert a test album
	_, err := db.Exec(`INSERT INTO albums (id, name, release_date, genre, price, description) VALUES (1, 'Old Album', '2022-01-01', 'Rock', 200, 'An old album')`)
	if err != nil {
		t.Fatalf("failed to insert test album: %v", err)
	}

	// Update the album
	album := &models.Album{ID: 1, Name: "Updated Album", ReleaseDate: "2022-01-01", Genre: "Rock", Price: 220, Description: "An updated album"}
	err = repo.UpdateAlbum(album)
	if err != nil {
		t.Fatalf("failed to update album: %v", err)
	}

	// Verify the update
	var updatedName string
	err = db.QueryRow("SELECT name FROM albums WHERE id = 1").Scan(&updatedName)
	if err != nil {
		t.Fatalf("failed to query updated album: %v", err)
	}

	if updatedName != "Updated Album" {
		t.Errorf("expected 'Updated Album', got '%s'", updatedName)
	}
}

func TestDeleteAlbum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	// Insert a test album
	_, err := db.Exec(`INSERT INTO albums (id, name, release_date, genre, price, description) VALUES (1, 'Test Album', '2022-01-01', 'Rock', 200, 'A test album')`)
	if err != nil {
		t.Fatalf("failed to insert test album: %v", err)
	}

	// Delete the album
	err = repo.DeleteAlbum(1)
	if err != nil {
		t.Fatalf("failed to delete album: %v", err)
	}

	// Verify the deletion
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM albums WHERE id = 1").Scan(&count)
	if err != nil {
		t.Fatalf("failed to query albums: %v", err)
	}

	if count != 0 {
		t.Errorf("expected 0 albums, got %d", count)
	}
}

func TestLinkMusiciansToAlbum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	// Insert a test album
	_, err := db.Exec(`INSERT INTO albums (id, name, release_date, genre, price, description) VALUES (1, 'Test Album', '2022-01-01', 'Rock', 200, 'A test album')`)
	if err != nil {
		t.Fatalf("failed to insert test album: %v", err)
	}

	// Test LinkMusiciansToAlbum
	err = repo.LinkMusiciansToAlbum(1, []uint{101, 102})
	if err != nil {
		t.Fatalf("failed to link musicians to album: %v", err)
	}

	// Verify musicians were linked to the album
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM album_musicians WHERE album_id = 1").Scan(&count)
	if err != nil {
		t.Fatalf("failed to count album_musicians: %v", err)
	}

	if count != 2 {
		t.Errorf("expected 2 musicians linked to album, got %d", count)
	}
}

func TestGetAlbumsByMusician(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := AlbumRepository{DB: db}

	// Insert test albums and musicians
	_, err := db.Exec(`
		INSERT INTO albums (id, name, release_date, genre, price, description) VALUES 
		(1, 'Rock Album', '2022-01-01', 'Rock', 200, 'First Album'),
		(2, 'Pop Album', '2022-02-01', 'Pop', 300, 'Second Album');
		INSERT INTO album_musicians (album_id, musician_id) VALUES 
		(1, 101),
		(2, 101),
		(1, 102);
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Test GetAlbumsByMusician
	albums, err := repo.GetAlbumsByMusician(101)
	if err != nil {
		t.Fatalf("failed to get albums by musician: %v", err)
	}

	// Verify the number of albums retrieved
	if len(albums) != 2 {
		t.Errorf("expected 2 albums for musician 101, got %d", len(albums))
	}

	// Check that the album names are correct
	if albums[0].Name != "Rock Album" || albums[1].Name != "Pop Album" {
		t.Errorf("unexpected album names for musician 101: %v, %v", albums[0].Name, albums[1].Name)
	}
}
