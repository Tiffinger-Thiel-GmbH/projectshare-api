package project

import (
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/api/dto"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler"
)

// Handler does everything which has to do with projects.
type Handler struct {
	ProjectRepository  handler.ProjectRepository
	DocumentRepository handler.DocumentRepository
}

// GetProjects loads a list of all available projects.
func (h Handler) GetProjects() ([]dto.Project, error) {
	projects, err := h.ProjectRepository.GetProjects()

	if err != nil {
		return nil, err
	}

	result := make([]dto.Project, len(projects))
	for i, project := range projects {
		result[i] = dto.Project{
			ID:   project.ID.String(),
			Name: project.Name,
		}
	}
	return result, nil
}

// GetProjectDocumentList loads all document metadata of a project.
func (h Handler) GetProjectDocumentList(projectID string) ([]dto.DocumentMetadata, error) {
	documents, err := h.DocumentRepository.GetDocumentsMetadata(projectID)

	if err != nil {
		return nil, err
	}

	result := make([]dto.DocumentMetadata, len(documents))
	for i, document := range documents {
		result[i] = dto.DocumentMetadata{
			ID:       document.ID.String(),
			Location: document.Location,
			Name:     document.Name,
			MimeType: document.MimeType,
		}
	}

	return result, nil
}

// AddProject creates a new project with a specific name.
// Note that currently the name does not have to be unique.
func (h Handler) AddProject(name string) (dto.Project, error) {
	document, err := h.ProjectRepository.AddProject(name)

	if err != nil {
		return dto.Project{}, err
	}

	return dto.Project{
		ID:   document.ID.String(),
		Name: document.Name,
	}, nil
}
