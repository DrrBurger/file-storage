package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"file-storage/internal/entities"
)

type FileStorage struct {
	Directory string
}

func NewFileStorage(dir string) *FileStorage {
	return &FileStorage{Directory: dir}
}

func (fs *FileStorage) Store(filePart entities.FilePart) error {
	filename := fmt.Sprintf("%s_part%d", filePart.FileID, filePart.PartID)
	path := filepath.Join(fs.Directory, filename)

	err := os.WriteFile(path, filePart.Data, 0644)
	return err
}

func (fs *FileStorage) Retrieve(fileID string, partID int) (entities.FilePart, error) {
	filename := fmt.Sprintf("%s_part%d", fileID, partID)
	path := filepath.Join(fs.Directory, filename)

	data, err := os.ReadFile(path)
	if err != nil {
		return entities.FilePart{}, err
	}
	return entities.FilePart{
		FileID: fileID,
		PartID: partID,
		Data:   data,
	}, nil
}
