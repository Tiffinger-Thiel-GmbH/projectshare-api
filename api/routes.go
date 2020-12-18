package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tiffinger-Thiel-GmbH/projectshare-api/api/dto"
	"github.com/go-chi/chi"
)

func (s *Server) writeJSON(w http.ResponseWriter, data interface{}) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(dataJSON)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) setupRoutes() {
	s.router.Get("/project", func(w http.ResponseWriter, r *http.Request) {
		projects, err := s.projectHandler.GetProjects()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.writeJSON(w, projects)
	})

	s.router.Get("/project/{projectID:[a-f0-9-]{36}}/document", func(w http.ResponseWriter, r *http.Request) {
		projectID := chi.URLParam(r, "projectID")
		// TODO: check for valid uuid?

		documents, err := s.projectHandler.GetProjectDocumentList(projectID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.writeJSON(w, documents)
	})

	s.router.Post("/project", func(w http.ResponseWriter, r *http.Request) {
		var p dto.CreateProject

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newProject, err := s.projectHandler.AddProject(p.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.writeJSON(w, newProject)
	})

	s.router.Get("/document/{location}/{documentID}", func(w http.ResponseWriter, r *http.Request) {
		location := chi.URLParam(r, "location")
		documentID := chi.URLParam(r, "documentID")

		documentData, metadata, err := s.documentHandler.GetDocument(location, documentID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", metadata.MimeType)
		w.Header().Add("Content-Disposition", "attachment; filename=\""+metadata.Name+"\"")

		_, err = w.Write(documentData)
		if err != nil {
			log.Println(err)
		}
	})

	s.router.Post("/document/{location}", func(w http.ResponseWriter, r *http.Request) {
		location := chi.URLParam(r, "location")

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		documentMetadata, err := s.documentHandler.AddDocument(location, file, handler.Size, handler.Filename, handler.Header.Get("Content-Type"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s.writeJSON(w, documentMetadata)
	})
}
