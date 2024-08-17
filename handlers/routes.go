package handlers

import (
	"net/http"
)

func SetRoutes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", addCorrectMIMEType(fs)))

	http.HandleFunc("/{$}", IndexHandler)
	http.HandleFunc("/editor", EditorHandler)
	http.HandleFunc("POST /publish", PublishHandler)
}
