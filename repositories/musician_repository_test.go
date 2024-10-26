package repositories

import (
	"jukebox/models"
	"testing"
)

func TestCreateMusician(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := MusicianRepository{DB: db}

	// Successful musician creation
	t.Run("successful create musician", func(t *testing.T) {
		musician := &models.Musician{Name: "John Doe", MusicianType: "Guitarist"}
		err := repo.CreateMusician(musician)
		if err != nil {
			t.Fatalf("failed to create musician: %v", err)
		}

		if musician.ID == 0 {
			t.Errorf("expected a valid musician ID, got 0")
		}
	})
}

func TestGetMusicians(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := MusicianRepository{DB: db}

	// Insert test musicians
	_, err := db.Exec(`
		INSERT INTO musicians (id, name, musician_type) VALUES 
		(1, 'John Doe', 'Guitarist'),
		(2, 'Jane Smith', 'Vocalist');
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Successful retrieval of musicians
	t.Run("successful get musicians", func(t *testing.T) {
		musicians, err := repo.GetMusicians()
		if err != nil {
			t.Fatalf("failed to get musicians: %v", err)
		}

		if len(musicians) != 2 {
			t.Errorf("expected 2 musicians, got %d", len(musicians))
		}
	})

	// Error scenario: No musicians found
	t.Run("no musicians found", func(t *testing.T) {
		_, err := db.Exec("DELETE FROM musicians")
		if err != nil {
			t.Fatalf("failed to clear musicians: %v", err)
		}

		musicians, err := repo.GetMusicians()
		if err != nil {
			t.Fatalf("failed to get musicians: %v", err)
		}

		if len(musicians) != 0 {
			t.Errorf("expected 0 musicians, got %d", len(musicians))
		}
	})
}

func TestUpdateMusician(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := MusicianRepository{DB: db}

	// Insert a test musician
	_, err := db.Exec("INSERT INTO musicians (id, name, musician_type) VALUES (1, 'John Doe', 'Guitarist')")
	if err != nil {
		t.Fatalf("failed to insert test musician: %v", err)
	}

	t.Run("successful update musician", func(t *testing.T) {
		musician := &models.Musician{ID: 1, Name: "Johnny Doe", MusicianType: "Bassist"}
		err := repo.UpdateMusician(musician)
		if err != nil {
			t.Fatalf("failed to update musician: %v", err)
		}
	})
}

func TestDeleteMusician(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := MusicianRepository{DB: db}

	// Insert a test musician
	_, err := db.Exec("INSERT INTO musicians (id, name, musician_type) VALUES (1, 'John Doe', 'Guitarist')")
	if err != nil {
		t.Fatalf("failed to insert test musician: %v", err)
	}

	// Successful musician deletion
	t.Run("successful delete musician", func(t *testing.T) {
		err := repo.DeleteMusician(1)
		if err != nil {
			t.Fatalf("failed to delete musician: %v", err)
		}
	})
}

func TestGetMusiciansByAlbum(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := MusicianRepository{DB: db}

	// Insert test albums and musicians
	_, err := db.Exec(`
		INSERT INTO albums (id, name, release_date, genre, price, description) VALUES 
		(1, 'Rock Album', '2022-01-01', 'Rock', 200, 'First Album');
		INSERT INTO musicians (id, name, musician_type) VALUES 
		(101, 'John Doe', 'Guitarist'),
		(102, 'Jane Smith', 'Vocalist');
		INSERT INTO album_musicians (album_id, musician_id) VALUES 
		(1, 101),
		(1, 102);
	`)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Successful get musicians by album
	t.Run("successful get musicians by album", func(t *testing.T) {
		musicians, err := repo.GetMusiciansByAlbum(1)
		if err != nil {
			t.Fatalf("failed to get musicians by album: %v", err)
		}

		if len(musicians) != 2 {
			t.Errorf("expected 2 musicians, got %d", len(musicians))
		}
	})

	// Error scenario: No musicians for album
	t.Run("no musicians found for album", func(t *testing.T) {
		musicians, err := repo.GetMusiciansByAlbum(999)
		if err != nil {
			t.Fatalf("failed to get musicians by album: %v", err)
		}

		if len(musicians) != 0 {
			t.Errorf("expected 0 musicians, got %d", len(musicians))
		}
	})
}
