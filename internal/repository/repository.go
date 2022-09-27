package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/storage-service/internal/domain"
)

const (
	fileTable = "files"
)

type Files interface {
	Create(file *domain.File) (uint, error)
	GetByID(ID int) (*domain.File, error)
	GetAllIds() []domain.FileTypeIds
	GetIds(fType string) []int
	DeleteOne(ID int) (int, error)
	Clear() error
}

type Repositories struct {
	Files Files
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Files: NewFileRepository(db),
	}
}
