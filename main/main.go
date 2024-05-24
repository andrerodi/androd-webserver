package main

import (
	"androd/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server at port 3333...")

	http.HandleFunc("/", handlers.NewHomeHandler().ServeHTTP)

	http.ListenAndServe(":3333", nil)
}
