package dto

// Project defines the data a project can have.
type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateProject defines the data needed for a new Project.
type CreateProject struct {
	Name string `json:"name"`
}
