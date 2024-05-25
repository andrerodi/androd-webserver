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

	id = normalizeId(id)

	log.Println("ID:", id)

	templ := templates.Projects(id, GetProjectsCount(), getProjectById(id))
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

func getProjectById(id int) templates.Project {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	projectsDir := filepath.Join(dir, "static", "projects")

	files, err := os.ReadDir(projectsDir)

	if err != nil {
		log.Fatal(err)
	}

	id = normalizeId(id)

	file := files[id-1]

	log.Println("Project ID:", id, "Name:", file.Name())

	contents, err := os.ReadFile(filepath.Join(projectsDir, file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	return templates.Project{
		Id:      id,
		Title:   file.Name(),
		Content: contents,
	}
}

func normalizeId(id int) int {
	if id < 1 {
		id = 1
	}

	if id > GetProjectsCount() {
		id = GetProjectsCount()
	}

	return id
}
