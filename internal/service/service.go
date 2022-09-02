package service

import (
	"mime/multipart"

	"github.com/onemgvv/storage-service/internal/repository"
	"github.com/onemgvv/storage-service/pkg/storage"
)

type Files interface {
	UploadFile(file *multipart.FileHeader) (uint, error)
	FindOne(id int) (string, error)
	Delete(id int) (int, error)
	Clear() error
}

type Service struct {
	Files Files
}

type Deps struct {
	Repos *repository.Repositories
	Storage *storage.Storage
}

func NewServices(deps *Deps) *Service {
	fileService := NewFileService(deps.Repos.Files, deps.Storage)
	return &Service{
		Files: fileService,
	}
}
