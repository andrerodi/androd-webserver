package handlers

import (
	"log"
	"net/http"
)

// GetStaticImgFS is a handler that serves static images from the filesystem.
type GetStaticImgFS struct{}

// NewGetStaticImgFS creates a new GetStaticImgFS handler.
func NewGetStaticImgFS() *GetStaticImgFS {
	return &GetStaticImgFS{}
}

// ServeHTTP serves static images from the filesystem.
func (handler *GetStaticImgFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving static image: ", r.URL.Path)
	imgfs := http.FileServer(http.Dir("./static/img"))
	http.Handle("/static/img/", http.StripPrefix("/static/img/", imgfs))
}
