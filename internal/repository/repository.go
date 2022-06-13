package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/storage-service.git/internal/domain"
)

const (
	fileTable = "files"
)

type Files interface {
	Create(file *domain.File) (string, error)
	GetByID(ID int) (*domain.File, error)
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
