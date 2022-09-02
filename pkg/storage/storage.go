package storage

import (
	"log"
	"os"
	"path/filepath"
)

type Storage struct {
	Directory string
}

func NewStorage(dir string) *Storage {
	return &Storage{
		Directory: dir,
	}
}

func (s *Storage) CreateDir(folder string) (string, error) {
	/**
	 *	make folder path and check exists same folder in path
	 *	if same folder exist in path return error with current message
	 */
	folderPath := filepath.Join(s.Directory, folder)
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		return "Folder already exist", nil
	}
	// If folder not exists create them in path and if something was wrong with creating return an error
	if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
		return "", err
	}
	// if folder created successfully return path to this folder
	return folderPath, nil
}

func (s *Storage) RemoveDir(folder string) (string, error) {
	/**
	 *	make folder path and check exists same folder in path
	 *	if same folder not exist in path return error with current message
	 */
	folderPath := filepath.Join(s.Directory, folder)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return "Folder not found", nil
	}
	// If folder exists in path remove them. If something was wrong with deleting file return an error
	if err := os.Remove(folderPath); err != nil {
		return "", err
	}
	// if folder successfully deleted return message
	return "Folder deleted successfully", nil
}

func (s *Storage) SaveFile(loadPath string, data []byte) error {
	/**
	 *	check exists file with same name in path
	 *	if file with same name exist in path return error with current message
	 */
	if _, err := os.Stat(loadPath); !os.IsNotExist(err) {
		return err
	}

	// Creating file in path
	newFile, err := os.Create(loadPath)
	if err != nil {
		return err
	}

	// close file after return from function
	defer func(newFile *os.File) {
		if err = newFile.Close(); err != nil {
			log.Println(err)
		}
	}(newFile)

	// write data in newly created file
	if _, err = newFile.Write(data); err != nil {
		return err
	}

	return nil
}

func (s *Storage) ReadFile(folder, file string) ([]byte, error) {
	/**
	 * Make Path to need file and if file not exist in path return empty string in bytes
	 */
	filePath := filepath.Join(s.Directory, folder, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []byte(""), nil
	}
	// if file exists read them
	fileData, err := os.ReadFile(filePath)
	// if something was wrong return empty byte string and error
	if err != nil {
		return []byte(""), err
	}
	// if all was ok return file data and nil like error
	return fileData, nil
}

func (s *Storage) DeleteFile(folder, file string) error {
	/**
	 * Make Path to need file and if file not exist in path return empty string
	 */
	filePath := filepath.Join(s.Directory, folder, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	// If file exists delete them, and if something was wrong with deleting return error
	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}
