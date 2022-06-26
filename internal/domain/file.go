package domain

import "time"

type FileType string

const (
	Text   FileType = "text"
	Config          = "config"
	Audio           = "audio"
	Image           = "image"
	Video           = "video"
)

type File struct {
	Id        int       `json:"-" db:"id"`
	Path      string    `json:"path" binding:"required"`
	Type      FileType  `json:"type" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
