package handlers

import (
	"androd/templates"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	about := templates.About()
	err := templates.Layout(about, "home").Render(r.Context(), w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
