package handlers

import (
	"androd/templates"
	"net/http"
)

// NewAboutHandler returns a new instance of the AboutHandler
func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

// AboutHandler handles the about page
type AboutHandler struct{}

// ServeHTTP serves the about page
func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ := templates.About()
	templates.Layout(templ, "about").Render(r.Context(), w)
}
