package model

import "github.com/google/uuid"

// Project is the model for Projects.
type Project struct {
	ID   uuid.UUID
	Name string
}
