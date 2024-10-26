package models

type Musician struct {
    ID           uint   `json:"id"`
    Name         string `json:"name"`
    MusicianType string `json:"musician_type"`
}
