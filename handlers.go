package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
