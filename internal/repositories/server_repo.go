package repositories

import (
	"file-storage/internal/entities"
)

type StorageServer interface {
	Store(filePart entities.FilePart) error
	Retrieve(fileID string, partID int) (entities.FilePart, error)
}

type ServerRepository struct {
	servers []StorageServer
}

func NewServerRepository(servers []StorageServer) *ServerRepository {
	return &ServerRepository{servers: servers}
}

func (r *ServerRepository) StoreFileParts(parts []entities.FilePart) error {
	for i, part := range parts {
		serverIndex := i % len(r.servers)
		err := r.servers[serverIndex].Store(part)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ServerRepository) RetrieveFileParts(fileID string, partsCount int) ([]entities.FilePart, error) {
	var parts []entities.FilePart
	for i := 0; i < partsCount; i++ {
		serverIndex := i % len(r.servers)
		part, err := r.servers[serverIndex].Retrieve(fileID, i)
		if err != nil {
			return nil, err
		}
		parts = append(parts, part)
	}
	return parts, nil
}
