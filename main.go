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
	fmt.Fprint(w, "<html><head><title>Default Page</title></head><body><h1>Welcome to the Redirector Service</h1><p>Use /sa to go to Site A or /sc to go to Site C.</p></body></html>")
}

func main() {
	http.HandleFunc("/", redirectHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve 404 page for unmatched routes
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<html><head><title>404 Not Found</title></head><body><h1>404 Not Found</h1><p>The page you're looking for does not exist.</p></body></html>")
	})

	log.Println("Server is running on port 443...")
	if err := http.ListenAndServeTLS(":443", "", "", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
