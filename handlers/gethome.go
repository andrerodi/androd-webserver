package handlers

import (
	"androd/templates"

	"github.com/labstack/echo/v4"
)

// Path: handlers/getabout.go

type echoHomeHandler struct{}

func newEchoHomeHandler() *echoHomeHandler {
	return &echoHomeHandler{}
}

func (handler *echoHomeHandler) ServeHTTP(c echo.Context) error {
	templ := templates.Home()
	err := templates.Layout(templ, "home").Render(c.Request().Context(), c.Response())

	if err != nil {
		return err
	}

	return nil
}

func (h *AggregateHandler) Home() {
	h.Echo.GET("/", newEchoHomeHandler().ServeHTTP)
}
