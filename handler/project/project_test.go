package project_test

import (
	"testing"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/api/dto"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/model"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/project"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/repository/memory"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/test"
	"github.com/google/uuid"
)

func NewMockedProjectRepo() *memory.ProjectRepository {
	p := memory.ProjectRepository{
		Projects: []model.Project{
			{
				ID:   uuid.MustParse("0975ea14-56a0-4d40-940c-c9e94aa6b359"),
				Name: "ProjectX",
			},
			{
				ID:   uuid.MustParse("22ab3304-5fd8-45ca-b7bd-0ae9836f4e28"),
				Name: "ProjectY",
			},
		},
	}
	return &p
}

func TestHandler_GetProjects(t *testing.T) {
	tests := []struct {
		name string
		h    project.Handler
		want []dto.Project
	}{
		{
			name: "simple GetProjects",
			h: project.Handler{
				ProjectRepository: NewMockedProjectRepo(),
			},
			want: []dto.Project{
				{
					ID:   "0975ea14-56a0-4d40-940c-c9e94aa6b359",
					Name: "ProjectX",
				},
				{
					ID:   "22ab3304-5fd8-45ca-b7bd-0ae9836f4e28",
					Name: "ProjectY",
				},
			},
		},
		{
			name: "empty slice",
			h: project.Handler{
				ProjectRepository: &memory.ProjectRepository{},
			},
			want: []dto.Project{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetProjects()
			test.ExpectedError(t, nil, err)
			test.Equals(t, tt.want, got)
		})
	}
}

func TestHandler_AddProject(t *testing.T) {
	tests := []struct {
		name        string
		h           project.Handler
		projectName string
	}{
		{
			name: "add simple project",
			h: project.Handler{
				ProjectRepository: NewMockedProjectRepo(),
			},
			projectName: "newProject",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.AddProject(tt.projectName)
			if test.ExpectedError(t, nil, err) != test.IsCorrectNil {
				return
			}
			test.Equals(t, tt.projectName, got.Name)
			test.Assert(t, got.ID != "", "ID has to be set")
		})
	}
}

func TestHandler_GetProjectDocumentList(t *testing.T) {

	var text1 = []byte("This is a file")
	var meta1 = model.Metadata{
		ID:       uuid.MustParse("270b607a-ba79-4a49-86fd-d43ca5ebd3cb"),
		Location: "0975ea14-56a0-4d40-940c-c9e94aa6b359",
		MimeType: "text/plain",
		Name:     "Superfile.txt",
		Size:     int64(len(text1)),
	}
	var documents = []model.Document{
		{
			Metadata: meta1,
			Data:     text1,
		},
	}

	simpleRepo := memory.DocumentRepository{
		Files: documents,
	}

	tests := []struct {
		name      string
		h         project.Handler
		projectID string
		want      []dto.DocumentMetadata
	}{
		{
			name: "get all documents of a project",
			h: project.Handler{
				DocumentRepository: &simpleRepo,
			},
			projectID: "0975ea14-56a0-4d40-940c-c9e94aa6b359",
			want: []dto.DocumentMetadata{
				{
					ID:       meta1.ID.String(),
					Location: meta1.Location,
					MimeType: meta1.MimeType,
					Name:     meta1.Name,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetProjectDocumentList(tt.projectID)
			if test.ExpectedError(t, nil, err) != test.IsCorrectNil {
				return
			}
			test.Equals(t, tt.want, got)
		})
	}
}
