package services

import (
	"jukebox/models"
)

// AlbumServiceInterface defines the methods that must be implemented by any album service.
type AlbumServiceInterface interface {
	CreateAlbum(album *models.Album) error
	UpdateAlbum(album *models.Album) error
	DeleteAlbum(albumID uint) error
	GetAlbums() ([]models.Album, error)
	LinkMusiciansToAlbum(albumID uint, musicianIDs []uint) error
	GetAlbumsByMusician(musicianID uint) ([]models.Album, error)
}

// MusicianServiceInterface defines the methods that must be implemented by any musician service.
type MusicianServiceInterface interface {
	CreateMusician(musician *models.Musician) error
	UpdateMusician(musician *models.Musician) error
	DeleteMusician(musicianID uint) error
	GetMusicians() ([]models.Musician, error)
	GetMusiciansByAlbum(albumID uint) ([]models.Musician, error)
}
