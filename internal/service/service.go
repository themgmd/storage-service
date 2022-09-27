package service

import (
	"github.com/patrickmn/go-cache"
	"mime/multipart"

	"github.com/onemgvv/storage-service/internal/domain"
	"github.com/onemgvv/storage-service/internal/repository"
	"github.com/onemgvv/storage-service/pkg/storage"
)

type Files interface {
	UploadFile(file *multipart.FileHeader) (uint, error)
	GetFile(id int, params domain.FileParams) ([]byte, error)
	AllFiles(fType string) map[string][]int
	Delete(id int) (int, error)
}

type Service struct {
	Files Files
}

type Deps struct {
	Repos   *repository.Repositories
	Storage *storage.Storage
	Cache   *cache.Cache
}

func NewServices(deps *Deps) *Service {
	fileService := NewFileService(deps.Repos.Files, deps.Storage)
	return &Service{
		Files: fileService,
	}
}
