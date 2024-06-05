package main

import (
	"androd/handlers"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Starting server at port 3333...")

	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/static", "static")

	handlers.InitHandlers(e)
	e.Logger.Fatal(e.Start(":3333"))
}
