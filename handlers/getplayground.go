package handlers

import (
	"androd/templates"
	"net/http"
)

// NewPlaygroundHandler returns a new instance of the PlaygroundHandler
func NewPlaygroundHandler() *PlaygroundHandler {
	return &PlaygroundHandler{}
}

// PlaygroundHandler handles the playground page
type PlaygroundHandler struct{}

// ServeHTTP serves the playground page
func (h *PlaygroundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ := templates.Playground()
	templates.Layout(templ, "playground").Render(r.Context(), w)
}

// Path: templates/playground.templ
