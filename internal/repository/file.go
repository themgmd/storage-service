package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/onemgvv/storage-service/internal/domain"
	"log"
)

type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) Create(file *domain.File) (uint, error) {
	var id uint
	query := fmt.Sprintf("INSERT INTO %s (path, type, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", fileTable)
	row := r.db.QueryRow(query, file.Path, file.Type, file.CreatedAt, file.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FileRepository) GetByID(ID int) (*domain.File, error) {
	var file domain.File

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", fileTable)
	if err := r.db.Get(&file, query, ID); err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) DeleteOne(ID int) (int, error) {
	var id int

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", fileTable)
	row := r.db.QueryRow(query, ID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *FileRepository) GetAllIds() []domain.FileTypeIds {
	var ft []domain.FileTypeIds
	query := fmt.Sprintf("SELECT id, type FROM %s", fileTable)

	err := r.db.Select(&ft, query)
	if err != nil {
		log.Println(err.Error())
	}

	return ft
}

func (r *FileRepository) GetIds(fType string) []int {
	var ids []int
	query := fmt.Sprintf("SELECT id FROM %s WHERE $1", fileTable)

	err := r.db.Select(&ids, query, fType)
	if err != nil {
		log.Println(err.Error())
	}

	return ids
}

func (r *FileRepository) Clear() error {
	query := fmt.Sprintf("TRUNCATE TABLE %s", fileTable)
	if _, err := r.db.Query(query); err != nil {
		return err
	}

	return nil
}
