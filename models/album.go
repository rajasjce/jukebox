package models

type Album struct {
    ID          uint   `json:"id"`
    Name        string `json:"name"`
    ReleaseDate string `json:"release_date"`
    Genre       string `json:"genre"`
    Price       float64 `json:"price"`
    Description string `json:"description"`
}
