package api

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/document"
	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/handler/project"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Errors returned if the usage of Server was wrong.
var (
	ErrNotInitialized     = errors.New("the server is not initialized, call Server.Init()")
	ErrAlreadyInitialized = errors.New("the server is already initialized")
)

// Server combines all handlers and starts the webservice using chi.
type Server struct {
	options         options
	projectHandler  project.Handler
	documentHandler document.Handler
	router          *chi.Mux
	isInitialized   bool
}

// Init sets up the whole server.
func (s *Server) Init() {
	if s.isInitialized {
		panic(ErrAlreadyInitialized)
	}

	// If the repos are still missing, use the default implementation: AWS
	if s.options.documentRepo == nil || s.options.projectRepo == nil {
		s.With(AWS("eu-west-1"))
	}

	s.projectHandler = project.Handler{
		ProjectRepository:  s.options.projectRepo,
		DocumentRepository: s.options.documentRepo,
	}

	s.documentHandler = document.Handler{
		DocumentRepository: s.options.documentRepo,
	}

	// Create router.
	s.router = chi.NewRouter()

	// Add middlewares.
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	s.router.Use(middleware.Logger)

	// Add routes.
	s.setupRoutes()
	s.isInitialized = true
}

// Serve finally starts the webserver.
// It is blocking always returns an error.
func (s *Server) Serve() error {
	if !s.isInitialized {
		return ErrNotInitialized
	}

	if s.options.port == 0 {
		s.options.port = 7777
	}

	host := net.JoinHostPort("", strconv.Itoa(s.options.port))
	log.Println("Serve ProjectShareAPI on " + host)
	err := http.ListenAndServe(host, s.router)
	return err
}
