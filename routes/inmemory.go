package routes

import "jukebox/models"

type InMemoryAlbumService struct {
	albums []models.Album
}

// InMemoryMusicianService is a mock implementation of MusicianServiceInterface for testing purposes.
type InMemoryMusicianService struct {
	musicians []models.Musician
}

func (s *InMemoryAlbumService) CreateAlbum(album *models.Album) error {
	album.ID = uint(len(s.albums) + 1) // Assign a new ID (for simplicity)
	s.albums = append(s.albums, *album)
	return nil
}

func (s *InMemoryAlbumService) LinkMusiciansToAlbum(albumID uint, musicianIDs []uint) error {
	// In-memory implementation to link musicians to an album (no-op for simplicity)
	return nil
}

func (s *InMemoryAlbumService) GetAlbums() ([]models.Album, error) {
	return s.albums, nil
}

func (s *InMemoryAlbumService) UpdateAlbum(album *models.Album) error {
	for i, a := range s.albums {
		if a.ID == album.ID {
			s.albums[i] = *album
			return nil
		}
	}
	return nil
}

func (s *InMemoryAlbumService) DeleteAlbum(albumID uint) error {
	for i, a := range s.albums {
		if a.ID == albumID {
			s.albums = append(s.albums[:i], s.albums[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *InMemoryAlbumService) GetAlbumsByMusician(musicianID uint) ([]models.Album, error) {
	// Return albums linked to the specified musician ID (for simplicity, return empty)
	return []models.Album{}, nil
}

func (s *InMemoryMusicianService) CreateMusician(musician *models.Musician) error {
	musician.ID = uint(len(s.musicians) + 1) // Assign a new ID (for simplicity)
	s.musicians = append(s.musicians, *musician)
	return nil
}

func (s *InMemoryMusicianService) GetMusicians() ([]models.Musician, error) {
	return s.musicians, nil
}

func (s *InMemoryMusicianService) UpdateMusician(musician *models.Musician) error {
	for i, m := range s.musicians {
		if m.ID == musician.ID {
			s.musicians[i] = *musician
			return nil
		}
	}
	return nil
}

func (s *InMemoryMusicianService) DeleteMusician(musicianID uint) error {
	for i, m := range s.musicians {
		if m.ID == musicianID {
			s.musicians = append(s.musicians[:i], s.musicians[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *InMemoryMusicianService) GetMusiciansByAlbum(albumID uint) ([]models.Musician, error) {
	// Return musicians linked to the specified album ID (for simplicity, return empty)
	return []models.Musician{}, nil
}
