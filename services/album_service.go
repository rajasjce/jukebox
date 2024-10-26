package services

import (
	"errors"
	"jukebox/models"
	"jukebox/repositories"
)

// AlbumService is the real implementation which uses the repository.
type AlbumService struct {
	Repo *repositories.AlbumRepository
}

// CreateAlbum validates and creates a new album.
func (s *AlbumService) CreateAlbum(album *models.Album) error {
	// Basic Validation
	if len(album.Name) < 5 {
		return errors.New("album name must be at least 5 characters long")
	}
	if album.Price < 100 || album.Price > 1000 {
		return errors.New("price must be between 100 and 1000")
	}
	return s.Repo.CreateAlbum(album)
}

// UpdateAlbum updates an existing album
func (s *AlbumService) UpdateAlbum(album *models.Album) error {
	return s.Repo.UpdateAlbum(album)
}

// GetAlbums retrieves all albums from the repository.
func (s *AlbumService) GetAlbums() ([]models.Album, error) {
	// Simply call the repository to get all albums
	return s.Repo.GetAlbums()
}

// DeleteAlbum deletes an album by ID
func (s *AlbumService) DeleteAlbum(albumID uint) error {
	return s.Repo.DeleteAlbum(albumID)
}

// GetAlbumsByMusician retrieves albums for a specific musician sorted by price
func (s *AlbumService) GetAlbumsByMusician(musicianID uint) ([]models.Album, error) {
	return s.Repo.GetAlbumsByMusician(musicianID)
}

func (s *AlbumService) LinkMusiciansToAlbum(albumID uint, musicianIDs []uint) error {
	return s.Repo.LinkMusiciansToAlbum(albumID, musicianIDs)
}
