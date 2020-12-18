package memory

import (
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/google/uuid"
)

// ProjectRepository implements the handler.ProjectRepository interface by using just a in-memory storage.
// Therefore all data gets lost if you restart the application.
type ProjectRepository struct {
	Projects []model.Project
}

// GetProjects returns the currently available projects.
func (p *ProjectRepository) GetProjects() ([]model.Project, error) {
	return p.Projects, nil
}

// AddProject creates a new project with the given name.
func (p *ProjectRepository) AddProject(name string) (model.Project, error) {
	p.Projects = append(p.Projects, model.Project{
		ID:   uuid.New(),
		Name: name,
	})
	return p.Projects[len(p.Projects)-1], nil
}
