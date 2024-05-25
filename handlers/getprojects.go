package handlers

import (
	"androd/templates"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// GetProjectsHandler handles the /projects route
type ProjectsHandler struct{}

// NewProjectsHandler creates a new GetProjectsHandler
func NewProjectsHandler() *ProjectsHandler {
	return &ProjectsHandler{}
}

// ServeHTTP serves the /projects route
func (h *ProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		id = 1
	}

	log.Println("ID:", id)

	templ := templates.Projects(id)
	templates.Layout(templ, "projects").Render(r.Context(), w)
}

func GetProjectsCount() int {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	projectsDir := filepath.Join(dir, "static", "projects")

	files, err := os.ReadDir(projectsDir)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Projects count:", len(files))

	return len(files)
}
