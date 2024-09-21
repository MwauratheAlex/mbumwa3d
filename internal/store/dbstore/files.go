package dbstore

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

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
		FileDir: "./public/model_storage",
	}
}

func (s *FileStore) SaveToDisk(file multipart.File, filename string) (string, error) {
	defer file.Close()
	if err := os.MkdirAll(s.FileDir, os.ModePerm); err != nil {
		return "", err
	}

	fileNameInDisk := s.generateUniqueFilename(filename)

	dstPath := filepath.Join(s.FileDir, fileNameInDisk)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}
	return fileNameInDisk, nil
}

func (s *FileStore) generateUniqueFilename(originalFileName string) string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%d-%s", timestamp, originalFileName)

}
