package main

import (
	"androd/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting server at port 3333...")

	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Build the path to the static directory
	staticDir := filepath.Join(dir, "static")
	log.Println("Serving static files from: ", staticDir)
	// Serve static images from the filesystem
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handlers.NewHomeHandler().ServeHTTP)
	http.HandleFunc("/about", handlers.NewAboutHandler().ServeHTTP)
	http.HandleFunc("/projects", handlers.NewProjectsHandler().ServeHTTP)
	http.HandleFunc("/playground", handlers.NewPlaygroundHandler().ServeHTTP)

	http.ListenAndServe(":3333", nil)
}
