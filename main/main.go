package main

import (
	"androd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fmt"
)

func main() {
	fmt.Println("Starting server at port 3333...")

	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/static", "static")

	// e.GET("/", handlers.NewEchoHomeHandler().ServeHTTP)
	// e.GET("/about", handlers.NewEchoAboutHandler().ServeHTTP)
	// e.GET("/projects", handlers.NewEchoProjectHandler().ServeHTTP)

	handlers.InitHandlers(e)
	e.Logger.Fatal(e.Start(":3333"))
}
