package main

import (
	"log"
	"net/http"

	"github.com/justinliang1020/html-surf/handlers"
	"github.com/justinliang1020/html-surf/services"
)

func main() {
	services.SeedSnippets()
	handlers.SetRoutes()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingMiddleware(http.DefaultServeMux)))
}
