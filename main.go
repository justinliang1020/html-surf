package main

import (
	"log"
	"net/http"
)

func main() {
	SeedSnippets()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", addCorrectMIMEType(fs)))
	http.HandleFunc("/editor", EditorHandler)
	http.HandleFunc("POST /publish", PublishHandler)
	http.HandleFunc("/{$}", IndexHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", LoggingMiddleware(http.DefaultServeMux)))
}
