package services

import (
	"errors"
	"jukebox/models"
	"jukebox/repositories"
)

type MusicianService struct {
	Repo *repositories.MusicianRepository
}

// CreateMusician validates and creates a new musician.
func (s *MusicianService) CreateMusician(musician *models.Musician) error {
	if len(musician.Name) < 3 {
		return errors.New("musician name must be at least 3 characters long")
	}
	return s.Repo.CreateMusician(musician)
}

// GetMusicians retrieves all musicians from the database.
func (s *MusicianService) GetMusicians() ([]models.Musician, error) {
	return s.Repo.GetMusicians()
}

// UpdateMusician updates an existing musician.
func (s *MusicianService) UpdateMusician(musician *models.Musician) error {
	return s.Repo.UpdateMusician(musician)
}

// DeleteMusician deletes a musician by ID.
func (s *MusicianService) DeleteMusician(musicianID uint) error {
	return s.Repo.DeleteMusician(musicianID)
}

// GetMusiciansByAlbum retrieves musicians for a specific album.
func (s *MusicianService) GetMusiciansByAlbum(albumID uint) ([]models.Musician, error) {
	return s.Repo.GetMusiciansByAlbum(albumID)
}
