package handlers

import (
	"androd/templates"

	"github.com/labstack/echo/v4"
)

// Echo handler for the playground page.

// newEchoPlaygroundHandler returns a new EchoPlaygroundHandler.
func newEchoPlaygroundHandler() *echoPlaygroundHandler {
	return &echoPlaygroundHandler{}
}

// echoPlaygroundHandler is a handler for the playground page.
type echoPlaygroundHandler struct{}

// ServeHTTP serves the playground page.
func (h *echoPlaygroundHandler) ServeHTTP(c echo.Context) error {
	templ := templates.Playground()

	if err := templ.Render(c.Request().Context(), c.Response()); err != nil {
		return err
	}

	return nil
}

// Handle all routes for the playground page.
// this *echo.Echo is a pointer to the echo instance.
// Extend our type by echo.Echo.

func (e *AggregateHandler) Playground() {
	e.Echo.GET("/playground", newEchoPlaygroundHandler().ServeHTTP)
}
