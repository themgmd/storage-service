package storage

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestStorage_CreateDir(t *testing.T) {
	var storage = NewStorage("/home")

	str, err := storage.CreateDir("images")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, str, "/home/images")
}

func TestStorage_RemoveDir(t *testing.T) {
	var storage = NewStorage("/home")

	str, err := storage.RemoveDir("images")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, str, "Folder deleted successfully")
}

func TestStorage_SaveFile(t *testing.T) {
	var storage = NewStorage("/home")
	err := storage.SaveFile(storage.Directory+"/file.txt", []byte("Test File Text"))
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, err, nil)
}

func TestStorage_ReadFile(t *testing.T) {
	var storage = NewStorage("/home")

	data, err := storage.ReadFile("", "file.txt")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, string(data), "Test File Text")
}

func TestStorage_DeleteFile(t *testing.T) {
	var storage = NewStorage("/home")

	err := storage.DeleteFile(storage.Directory + "/file.txt")
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, err, nil)
}
