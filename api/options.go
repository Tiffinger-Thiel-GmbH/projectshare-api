package api

// File options provides possible configuration options for the server.
// They can be passed to the NewServer function like this:
//
// s := api.Server{}
// s.With(api.InMemory())
//
// This is an often used Pattern (https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
// which can be used as a very flexible alternative to for example just a struct:
//
// (Pseudocode:)
// s := Server{
// 	Port: 7777,
//  ...
// }
//
// Which would be also a valid way for simple options.
//
// Also possible would be a
// func NewServer(port string)
// But this would not be backwards compatible as adding new parameters would break that.
// Also with a lot of options this would become soon very clumsy.
//
// So the functional options provides better extendability by still maintaining backwards compatibility and you can
// validate the options inside of the With-functions.

import (
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/repository/aws"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/repository/memory"
)

// With enables the passed options.
// This has to be called before Server.Init.
func (s *Server) With(options ...Option) *Server {
	if s.isInitialized {
		panic(ErrAlreadyInitialized)
	}

	for _, o := range options {
		o(s)
	}

	return s
}

// Option is basically just a function which accepts a server-pointer as param which it can change in any way.
type Option = func(s *Server)

type options struct {
	port         int
	projectRepo  handler.ProjectRepository
	documentRepo handler.DocumentRepository
}

// Port adds the port to the server on which it should serve the api.
func Port(port int) Option {
	return func(s *Server) {
		s.options.port = port
	}
}

// DocumentRepository allows passing a different implementation of the handler.DocumentRepository interface.
func DocumentRepository(documentRepo handler.DocumentRepository) Option {
	return func(s *Server) {
		s.options.documentRepo = documentRepo
	}
}

// ProjectRepository allows passing a different implementation of the handler.DocumentRepository interface.
func ProjectRepository(projectRepo handler.ProjectRepository) Option {
	return func(s *Server) {
		s.options.projectRepo = projectRepo
	}
}

// AWS enables the use of repositories based on AWS S3.
func AWS(region string) Option {
	return func(s *Server) {
		documentRepo := aws.NewDocumentRepository(region)
		projectRepo := aws.ProjectRepository{
			DocumentRepository: documentRepo,
		}

		s.With(ProjectRepository(&projectRepo))
		s.With(DocumentRepository(&documentRepo))
	}
}

// InMemory enables the use of repositories based on in memory storage.
// The data will be lost after restart, but noc external server is needed.
func InMemory() Option {
	return func(s *Server) {
		s.With(ProjectRepository(&memory.ProjectRepository{}))
		s.With(DocumentRepository(&memory.DocumentRepository{}))
	}
}
