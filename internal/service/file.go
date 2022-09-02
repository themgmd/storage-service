package service

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/onemgvv/storage-service/internal/domain"
	"github.com/onemgvv/storage-service/internal/repository"
	"github.com/onemgvv/storage-service/pkg/filetype"
	"github.com/onemgvv/storage-service/pkg/storage"
)

var folders = map[domain.FileType]string{
	domain.Image: "images",
	domain.Video: "videos",
	domain.Audio: "audios",
	domain.Text:  "other",
	domain.DOCS:  "documents",
}

type FileService struct {
	repo repository.Files
	storage *storage.Storage
}

func NewFileService(repo repository.Files, storage *storage.Storage) *FileService {
	return &FileService{
		repo: repo,
		storage: storage,
	}
}

func (f *FileService) UploadFile(file *multipart.FileHeader) (uint, error) {
	openedFile, err := file.Open()
	defer openedFile.Close()
	if err != nil {
		return 0, err
	}

	extension := filepath.Ext(file.Filename)
	newName := uuid.New().String() + extension
	fType := filetype.DetectType(extension)

	if fType == "" {
		return 0, errors.New("Unknown file type")
	}

	data, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return 0, err
	}

	var sb strings.Builder
	sb.WriteString(filepath.Join(f.storage.Directory, folders[fType], newName))

	loadPath := sb.String()

	err = f.storage.SaveFile(loadPath, data)
	if err != nil {
		return 0, err
	}

	id, err := f.repo.Create(&domain.File{Path: loadPath, Type: fType})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (f *FileService) FindOne(id int) (string, error) {
	file, err := f.repo.GetByID(id)
	if err != nil {
		return "", err
	}

	return file.Path, nil
}

func (f *FileService) Delete(id int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (f FileService) Clear() error {
	//TODO implement me
	panic("implement me")
}
