package entities

type File struct {
	ID   string
	Data []byte
}

type FilePart struct {
	FileID string
	PartID int
	Data   []byte
}
