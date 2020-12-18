package handler

import (
	"io"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/google/uuid"
)

// DocumentRepository describes all actions which can be done with documents.
type DocumentRepository interface {
	GetDocumentsMetadata(location string) ([]model.Metadata, error)
	GetDocument(location string, documentID string) (model.Document, error)
	AddDocument(reader io.Reader, metadata model.Metadata) (uuid.UUID, error)
}

// ProjectRepository describes all actions which can be done with projects.
type ProjectRepository interface {
	GetProjects() ([]model.Project, error)
	AddProject(name string) (model.Project, error)
}
