package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ping" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "pong")

}
