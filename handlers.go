package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Snippets []Snippet
	}
	data := PageData{
		Snippets: GetLatestSnippets(3),
	}
	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EditorHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "editor.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PublishHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		http.Error(w, "Error processing request", http.StatusBadRequest)
		return
	}

	content := r.Form.Get("content")

	AddSnippet(content, "Placeholder Title", "Placeholder Author")
	response := `<div>
        Content published successfully!
    </div>`
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}
