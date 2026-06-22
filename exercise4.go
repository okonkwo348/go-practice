package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	username := r.FormValue("username")
	language := r.FormValue("lauguage")

	if len(username) == 0 {
		http.Error(w, "Bad request", 400)
		return
	}

	if len(language) == 0 {
		http.Error(w, "Bad request", 400)
		return
	}

	if len(username) != 0 && len(language) != 0 {
		fmt.Fprintf(w, "Hello %v, your are coding in %v!", username, language)
	}
}
