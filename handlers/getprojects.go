package handlers

import (
	"androd/templates"
	"encoding/json"
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

func getDirEntries(filetype string) []os.DirEntry {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	projectsDir := filepath.Join(dir, "static", "projects")

	files, err := os.ReadDir(projectsDir)

	if err != nil {
		log.Fatal(err)
	}

	var mdFiles []os.DirEntry
	// Get files of type filetype
	for i := 0; i < len(files); i++ {
		if filepath.Ext(files[i].Name()) == filetype {
			mdFiles = append(mdFiles, files[i])
			log.Println("Project:", files[i].Name())
		}
	}

	return mdFiles
}

func GetProjectsCount() int {
	files := getDirEntries(".md")

	return len(files)
}

type metaJson struct {
	Title string `json:"title"`
	Date  string `json:"date"`
}

func getProjectById(id int) templates.Project {
	id = normalizeId(id)

	files := getDirEntries(".md")
	metas := getDirEntries(".json")
	log.Println("JSON Files:", metas)
	file := files[id-1]
	meta := metas[id-1]
	log.Println("JSON File:", meta)

	// Parse meta to metaJson struct
	metaFile, err := os.ReadFile(filepath.Join("static", "projects", meta.Name()))

	if err != nil {
		log.Fatal(err)
	}

	metaJson := metaJson{}
	err = json.Unmarshal(metaFile, &metaJson)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Project ID:", id, "Name:", file.Name())

	projectsDir := filepath.Join("static", "projects")
	contents, err := os.ReadFile(filepath.Join(projectsDir, file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	return templates.Project{
		Id:      id,
		Title:   metaJson.Title,
		Date:    metaJson.Date,
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
