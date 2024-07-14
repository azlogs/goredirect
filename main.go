package main

import (
	"fmt"
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
		defaultPage(w, r)
	}
}

func defaultPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<html><head><title>Redirect test</title></head><body><h1>Welcome to the Redirector Service</h1><p>Use /sa to go to Site A or /sc to go to Site C.</p></body></html>")
}

func main() {
	http.HandleFunc("/", redirectHandler)
	log.Println("Server is running on port 80...")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
