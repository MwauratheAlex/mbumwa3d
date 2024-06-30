package dbstore

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/mwaurathealex/mbumwa3d/internal/initializers"
	"gorm.io/gorm"
)

type FileStore struct {
	db      *gorm.DB
	FileDir string
}

func NewFileStore() *FileStore {
	return &FileStore{
		db:      initializers.DB,
		FileDir: "./model_storage",
	}
}

func (s *FileStore) SaveToDisk(file multipart.File, filename string) (string, error) {
	defer file.Close()
	if err := os.MkdirAll(s.FileDir, os.ModePerm); err != nil {
		return "Error creating directory:", err
	}

	dstPath := filepath.Join(s.FileDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "Error creating file on disk:", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "Error copying file:", err
	}
	return "", nil
}
