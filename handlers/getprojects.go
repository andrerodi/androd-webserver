package handlers

import (
	"androd/templates"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

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

// newEchoProjectHandler returns a new instance of the ProjectHandler
func newEchoProjectHandler() *projectHandler {
	return &projectHandler{}
}

// projectHandler handles the project page
type projectHandler struct{}

type project struct {
	Id int `query:"id"`
}

// ServeHTTP serves the project page
func (h *projectHandler) ServeHTTP(c echo.Context) error {
	var project project
	if err := c.Bind(&project); err != nil {
		project.Id = 1
	}

	log.Println("Project ID:", project.Id)

	id := project.Id
	id = normalizeId(id)

	log.Println("ID:", id)

	templ := templates.Projects(id, GetProjectsCount(), getProjectById(id))
	if err := templates.Layout(templ, "projects").Render(c.Request().Context(), c.Response()); err != nil {
		return err
	}

	return nil
}

func (h *AggregateHandler) Projects() {
	h.Echo.GET("/projects", newEchoProjectHandler().ServeHTTP)
}
