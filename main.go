package main

import (
	"log"
	"net/http"
	"strings"
)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasPrefix(path, "/sa"):
		http.Redirect(w, r, "https://whatsmyreferer.com/", http.StatusFound)
	case strings.HasPrefix(path, "/sc"):
		http.Redirect(w, r, "https://whatsmyreferer.com/", http.StatusFound)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", redirectHandler)
	log.Println("Server is running on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
