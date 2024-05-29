package handlers

import (
	"androd/templates"
	"log"

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
	if err := templates.Layout(templ, "playground").Render(c.Request().Context(), c.Response()); err != nil {
		return err
	}

	return nil
}

// Handle htmx post request.
func (h *echoPlaygroundHandler) HandlePost(c echo.Context) error {
	log.Println("Handling htmx post request...")

	if c, err := c.Response().Writer.Write([]byte("Hello from HTMX!")); err != nil {
		log.Println("Bytes written to response writer", c)
		return err
	}

	return nil
}

// Handle all routes for the playground page.
// this *echo.Echo is a pointer to the echo instance.
// Extend our type by echo.Echo.

func (e *AggregateHandler) Playground() {
	e.Echo.GET("/playground", newEchoPlaygroundHandler().ServeHTTP)
	e.Echo.POST("/htmx-post", newEchoPlaygroundHandler().HandlePost)
}
