package model

import "github.com/google/uuid"

// Metadata is model for the metadata of documents.
type Metadata struct {
	ID       uuid.UUID
	Location string
	Name     string
	Size     int64
	MimeType string
}

// Document is the model for a whole Document.
type Document struct {
	Metadata
	Data []byte
}
