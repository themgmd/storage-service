package service

import (
	"github.com/onemgvv/storage-service.git/internal/domain"
	"github.com/onemgvv/storage-service.git/internal/repository"
)

var folders = map[domain.FileType]string{
	domain.Image: "images",
	domain.Video: "videos",
}

type FileService struct {
	repo repository.Files
}

func NewFileService(repo repository.Files) *FileService {
	return &FileService{
		repo: repo,
	}
}

func (f FileService) UploadFile(file domain.File) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileService) FindOne(id int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileService) Delete(id int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileService) Clear() error {
	//TODO implement me
	panic("implement me")
}
