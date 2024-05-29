package handlers

import "github.com/labstack/echo/v4"

type AggregateHandler struct {
	Echo *echo.Echo
}

func InitHandlers(e *echo.Echo) {
	// This function is empty. It is used to make sure the handlers package is included in the build.
	h := &AggregateHandler{Echo: e}

	h.Home()
	h.Projects()
	h.About()
	h.Playground()
}
