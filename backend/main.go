package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Simple health check endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from the Go Backend! Listening on port 8080.")
	})

	fmt.Println("Starting Go Server on port 8080...")
	// Start  the HTTP server on port 8080
	//
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
