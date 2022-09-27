package service

import (
	"bytes"
	"errors"
	"image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/onemgvv/storage-service/internal/domain"
	"github.com/onemgvv/storage-service/internal/repository"
	"github.com/onemgvv/storage-service/pkg/filetype"
	"github.com/onemgvv/storage-service/pkg/resize"
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
	repo    repository.Files
	storage *storage.Storage
}

func NewFileService(repo repository.Files, storage *storage.Storage) *FileService {
	return &FileService{
		repo:    repo,
		storage: storage,
	}
}

func (f *FileService) UploadFile(file *multipart.FileHeader) (uint, error) {
	openedFile, err := file.Open()
	if err != nil {
		return 0, err
	}

	extension := filepath.Ext(file.Filename)
	newName := uuid.New().String() + extension
	fType := filetype.DetectType(extension)

	if fType == "FE" {
		return 0, errors.New("unknown file type")
	}

	data, err := io.ReadAll(openedFile)
	if oErr := openedFile.Close(); err != nil {
		return 0, oErr
	}
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

func (f *FileService) GetFile(id int, params domain.FileParams) ([]byte, error) {
	// Get file record from DB
	file, err := f.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// If file not image return it without changes
	if file.Type != domain.Image {
		return readFile(file.Path)
	}

	// If file is image but request was without params return pure image
	if params.Width == 0 && params.Height == 0 {
		return readFile(file.Path)
	}

	// Resize image with requested params
	newImg, err := resize.ChangeSize(file.Path, params.Width, params.Height)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, newImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (f *FileService) Delete(id int) (int, error) {
	// Get file record from DB
	file, err := f.repo.GetByID(id)
	if err != nil {
		return 0, err
	}

	// Delete file from storage
	if err := f.storage.DeleteFile(file.Path); err != nil {
		return 0, err
	}

	// Delete record from db
	return f.repo.DeleteOne(id)
}

func (f *FileService) AllFiles(fType string) map[string][]int {
	if fType == "" {
		counts := map[string][]int{
			string(domain.Image): {},
			string(domain.Video): {},
			string(domain.Audio): {},
			string(domain.DOCS):  {},
			string(domain.Text):  {},
		}

		for _, v := range f.repo.GetAllIds() {
			counts[string(v.Type)] = append(counts[string(v.Type)], v.Id)
		}

		return counts
	} else {
		return map[string][]int{fType: f.repo.GetIds(fType)}
	}
}

func readFile(readPath string) ([]byte, error) {
	return os.ReadFile(readPath)
}
