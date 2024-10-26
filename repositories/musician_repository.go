package repositories

import (
	"database/sql"
	"jukebox/models"
	"log"
)

type MusicianRepository struct {
	DB *sql.DB
}

// CreateMusician inserts a new musician into the database and returns the inserted musician with the correct ID.
func (r *MusicianRepository) CreateMusician(musician *models.Musician) error {
	// Check if the musicians table is empty by seeing if any row exists
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM musicians LIMIT 1)").Scan(&exists)
	if err != nil {
		log.Println("Error checking if table is empty:", err)
		return err
	}

	if !exists {
		// If the table is empty, explicitly set the musician ID to 1
		musician.ID = 1
		_, err = r.DB.Exec("INSERT INTO musicians (id, name, musician_type) VALUES (?, ?, ?)",
			musician.ID, musician.Name, musician.MusicianType)
		if err != nil {
			log.Println("Error inserting musician with ID 1:", err)
			return err
		}
	} else {
		// Insert the musician into the database (ID will be auto-generated)
		result, err := r.DB.Exec("INSERT INTO musicians (name, musician_type) VALUES (?, ?)", musician.Name, musician.MusicianType)
		if err != nil {
			log.Println("Error inserting musician:", err)
			return err
		}

		// Get the last inserted ID
		id, err := result.LastInsertId()
		if err != nil {
			log.Println("Error retrieving last insert ID:", err)
			return err
		}

		// Set the ID in the musician object
		musician.ID = uint(id)
	}

	log.Println("Musician created with ID:", musician.ID)

	return nil
}

// GetMusicians retrieves all musicians from the database.
func (r *MusicianRepository) GetMusicians() ([]models.Musician, error) {
	rows, err := r.DB.Query("SELECT id, name, musician_type FROM musicians")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musicians []models.Musician
	for rows.Next() {
		var musician models.Musician
		if err := rows.Scan(&musician.ID, &musician.Name, &musician.MusicianType); err != nil {
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return musicians, nil
}

// UpdateMusician updates an existing musician in the database.
func (r *MusicianRepository) UpdateMusician(musician *models.Musician) error {
	_, err := r.DB.Exec("UPDATE musicians SET name = ?, musician_type = ? WHERE id = ?", musician.Name, musician.MusicianType, musician.ID)
	return err
}

// DeleteMusician deletes a musician by ID.
func (r *MusicianRepository) DeleteMusician(id uint) error {
	_, err := r.DB.Exec("DELETE FROM musicians WHERE id = ?", id)
	return err
}

// GetMusiciansByAlbum retrieves musicians for a specific album sorted by musician name.
func (r *MusicianRepository) GetMusiciansByAlbum(albumID uint) ([]models.Musician, error) {
	rows, err := r.DB.Query("SELECT m.id, m.name, m.musician_type FROM musicians m "+
		"JOIN album_musicians am ON m.id = am.musician_id WHERE am.album_id = ? ORDER BY m.name ASC", albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musicians []models.Musician
	for rows.Next() {
		var musician models.Musician
		if err := rows.Scan(&musician.ID, &musician.Name, &musician.MusicianType); err != nil {
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return musicians, nil
}
