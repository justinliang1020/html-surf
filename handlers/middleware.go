package handlers

import (
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, r.Proto, duration)
	})
}

func addCorrectMIMEType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ext := filepath.Ext(r.URL.Path)

		if ext != "" {
			mimeType := mime.TypeByExtension(ext)
			if mimeType != "" {
				w.Header().Set("Content-Type", mimeType)
			}
		}

		next.ServeHTTP(w, r)
	})
}
