package usecases

import (
	"errors"
	"file-storage/internal/entities"
)

func SplitFileIntoParts(file entities.File, partsCount int) ([]entities.FilePart, error) {
	if partsCount == 0 {
		return nil, errors.New("invalid parts count")
	}

	if partsCount < 0 {
		partsCount = -partsCount
	}

	var parts []entities.FilePart
	partSize := len(file.Data) / partsCount
	for i := 0; i < partsCount; i++ {
		start := i * partSize
		end := start + partSize
		if i == partsCount-1 {
			end = len(file.Data)
		}
		parts = append(parts, entities.FilePart{
			FileID: file.ID,
			PartID: i,
			Data:   file.Data[start:end],
		})
	}
	return parts, nil
}

func CombineFileFromParts(parts []entities.FilePart) (entities.File, error) {
	if len(parts) == 0 {
		return entities.File{}, errors.New("no parts provided")
	}

	var data []byte
	for _, part := range parts {
		data = append(data, part.Data...)
	}
	return entities.File{
		ID:   parts[0].FileID,
		Data: data,
	}, nil
}
