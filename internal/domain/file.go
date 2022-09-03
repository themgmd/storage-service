package domain

import "time"

type FileType string

const (
	Text   FileType = "text"
	Audio  FileType = "audio"
	Image  FileType = "image"
	Video  FileType = "video"
	DOCS   FileType = "docs"
)

type FileParams struct {
	Width  uint    `json:"width" form:"width"`
	Height uint    `json:"height" form:"height"`
}

type File struct {
	Id        int       `json:"-" db:"id"`
	Path      string    `json:"path" binding:"required"`
	Type      FileType  `json:"type" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
