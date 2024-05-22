package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server at port 3333...")

	http.HandleFunc("/", landingPageHandler)

	http.ListenAndServe(":3333", nil)
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
