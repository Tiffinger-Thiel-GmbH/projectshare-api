package document

import (
	"errors"
	"io"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/api/dto"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/google/uuid"
)

// Errors which can be returned by the Handler.
var (
	ErrWhileGettingFileStats = errors.New("error while getting file stats")
)

// Handler does everything which has to do with loading and saving documents.
type Handler struct {
	DocumentRepository handler.DocumentRepository
}

// GetDocument just loads a document by id and location and returns the content of it.
func (d Handler) GetDocument(location string, documentID string) ([]byte, dto.DocumentMetadata, error) {
	metadata, err := d.DocumentRepository.GetDocument(location, documentID)
	if err != nil {
		return nil, dto.DocumentMetadata{}, err
	}
	return metadata.Data, dto.DocumentMetadata{
		ID:       metadata.ID.String(),
		Location: metadata.Location,
		Name:     metadata.Name,
		MimeType: metadata.MimeType,
	}, nil
}

// AddDocument saves the given file at the given location.
// It returns the complete metadata including the new documentID.
func (d Handler) AddDocument(location string, reader io.Reader, size int64, fileName string, mimeType string) (dto.DocumentMetadata, error) {
	metadata := model.Metadata{
		ID:       uuid.UUID{},
		Location: location,
		Name:     fileName,
		Size:     size,
		MimeType: mimeType,
	}

	documentID, err := d.DocumentRepository.AddDocument(reader, metadata)

	if err != nil {
		return dto.DocumentMetadata{}, err
	}

	return dto.DocumentMetadata{
		ID:       documentID.String(),
		Location: metadata.Location,
		Name:     metadata.Name,
		MimeType: metadata.MimeType,
	}, err
}
