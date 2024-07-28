package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static", addCorrectMIMEType(fs)))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/editor", EditorHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", LoggingMiddleware(http.DefaultServeMux)))
}
