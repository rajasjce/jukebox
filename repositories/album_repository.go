package repositories

import (
	"database/sql"
	"jukebox/models"
)

type AlbumRepository struct {
	DB *sql.DB
}

// CreateAlbum inserts a new album into the database and returns the inserted album with the correct ID.
func (r *AlbumRepository) CreateAlbum(album *models.Album) error {
	// Check if the albums table is empty by seeing if any row exists
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM albums LIMIT 1)").Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		// If the table is empty, explicitly set the album ID to 1
		album.ID = 1
		_, err = r.DB.Exec("INSERT INTO albums (id, name, release_date, genre, price, description) VALUES (?, ?, ?, ?, ?, ?)",
			album.ID, album.Name, album.ReleaseDate, album.Genre, album.Price, album.Description)
		if err != nil {
			return err
		}
	} else {
		// Insert the album into the database (ID will be auto-generated)
		result, err := r.DB.Exec("INSERT INTO albums (name, release_date, genre, price, description) VALUES (?, ?, ?, ?, ?)",
			album.Name, album.ReleaseDate, album.Genre, album.Price, album.Description)
		if err != nil {
			return err
		}

		// Get the last inserted ID
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// Set the ID in the album object
		album.ID = uint(id)
	}

	return nil
}

// GetAlbums retrieves all albums from the database.
func (r *AlbumRepository) GetAlbums() ([]models.Album, error) {
	rows, err := r.DB.Query("SELECT id, name, release_date, genre, price, description FROM albums ORDER BY release_date ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Name, &album.ReleaseDate, &album.Genre, &album.Price, &album.Description); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

// UpdateAlbum updates an existing album in the database.
func (r *AlbumRepository) UpdateAlbum(album *models.Album) error {
	_, err := r.DB.Exec("UPDATE albums SET name = ?, release_date = ?, genre = ?, price = ?, description = ? WHERE id = ?",
		album.Name, album.ReleaseDate, album.Genre, album.Price, album.Description, album.ID)
	return err
}

// DeleteAlbum deletes an album by ID.
func (r *AlbumRepository) DeleteAlbum(id uint) error {
	_, err := r.DB.Exec("DELETE FROM albums WHERE id = ?", id)
	return err
}

func (r *AlbumRepository) LinkMusiciansToAlbum(albumID uint, musicianIDs []uint) error {
	// Insert each musician into album_musicians table
	for _, musicianID := range musicianIDs {
		_, err := r.DB.Exec("INSERT INTO album_musicians (album_id, musician_id) VALUES (?, ?)", albumID, musicianID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *AlbumRepository) GetAlbumsByMusician(musicianID uint) ([]models.Album, error) {
	rows, err := r.DB.Query(`
        SELECT a.id, a.name, a.release_date, a.genre, a.price, a.description
        FROM albums a
        JOIN album_musicians am ON a.id = am.album_id
        WHERE am.musician_id = ?
        ORDER BY a.price ASC
    `, musicianID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Name, &album.ReleaseDate, &album.Genre, &album.Price, &album.Description); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}
