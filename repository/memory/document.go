package memory

import (
	"errors"
	"fmt"
	"io"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/google/uuid"
)

// Errors which can be returned by DocumentRepository
var (
	ErrFileNotFound    = errors.New("file not found")
	ErrReadingNotFound = errors.New("error while reading stream")
)

// DocumentRepository implements the handler.DocumentRepository interface by using just a in-memory storage.
// Therefore all data gets lost if you restart the application.
type DocumentRepository struct {
	Files []model.Document
}

// GetDocumentsMetadata provides a list of all currently available documents with their metadata for a specific location.
func (r *DocumentRepository) GetDocumentsMetadata(location string) ([]model.Metadata, error) {
	var result []model.Metadata

	for _, d := range r.Files {
		if d.Location == location {
			result = append(result, d.Metadata)
		}
	}

	return result, nil
}

// GetDocument provides one specific document including its binary data.
func (r *DocumentRepository) GetDocument(location string, documentID string) (model.Document, error) {
	for _, d := range r.Files {
		if d.Location == location && d.ID.String() == documentID {
			return d, nil
		}
	}

	return model.Document{}, ErrFileNotFound
}

// AddDocument creates a new document from the given data.
// You have to provide the full metadata of the file.
// The field metadata.Size has to match the size of the document to be read!
// It returns the new id for the document. The field metadata.ID gets overwritten by this id.
func (r *DocumentRepository) AddDocument(document io.Reader, metadata model.Metadata) (uuid.UUID, error) {
	data := make([]byte, metadata.Size)
	_, err := document.Read(data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w:\n%v", ErrReadingNotFound, err)
	}

	metadata.ID = uuid.New()

	r.Files = append(r.Files, model.Document{
		Metadata: metadata,
		Data:     data,
	})
	return metadata.ID, nil
}
