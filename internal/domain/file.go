package domain

import "time"

type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
)

type File struct {
	Id        int       `json:"-" db:"id"`
	Path      string    `json:"path" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
