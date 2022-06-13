package service

import (
	"github.com/onemgvv/storage-service.git/internal/domain"
	"github.com/onemgvv/storage-service.git/internal/repository"
)

type Files interface {
	UploadFile(file domain.File) (string, error)
	FindOne(id int) (string, error)
	Delete(id int) (int, error)
	Clear() error
}

type Services struct {
	Files Files
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps *Deps) *Services {
	fileService := NewFileService(deps.Repos.Files)
	return &Services{
		Files: fileService,
	}
}
