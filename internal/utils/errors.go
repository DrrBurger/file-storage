package utils

type StorageError struct {
	Message string
}

func (e *StorageError) Error() string {
	return e.Message
}

func NewStorageError(msg string) error {
	return &StorageError{Message: msg}
}
