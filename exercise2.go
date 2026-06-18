package main

import (
	"fmt"
	"net/http"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest!"
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	fullName := fmt.Sprintf("Hello, %s!", name)

	fmt.Fprint(w, fullName)

}
