package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// As "/" acts as a catch-all, we need to actively check
	// to see if the path matched exactly.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home of Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)
	// IDs start at 1.
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Creating a snippet, huh?"))
}
