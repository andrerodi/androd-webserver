package handlers

import (
	"androd/templates"

	"github.com/labstack/echo/v4"
)

// NewAboutHandler returns a new instance of the AboutHandler
func newEchoAboutHandler() *aboutHandler {
	return &aboutHandler{}
}

// aboutHandler handles the about page
type aboutHandler struct{}

// ServeHTTP serves the about page
func (h *aboutHandler) ServeHTTP(c echo.Context) error {
	templ := templates.About()
	if err := templ.Render(c.Request().Context(), c.Response()); err != nil {
		return err
	}

	return nil
}

func (h *AggregateHandler) About() {
	h.Echo.GET("/about", newEchoAboutHandler().ServeHTTP)
}
