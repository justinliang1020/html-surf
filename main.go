package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", addCorrectMIMEType(fs)))
	http.HandleFunc("/editor", EditorHandler)
	http.HandleFunc("/publish", PublishHandler)
	http.HandleFunc("/{$}", IndexHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", LoggingMiddleware(http.DefaultServeMux)))
}
