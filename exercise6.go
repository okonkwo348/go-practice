package main

import (
	"fmt"

	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/dashboard" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	head := r.Header.Get("X-API-Key")
	g := "secret123"

	if head != g {
		http.Error(w, "User Unauthorized", http.StatusUnauthorized)
		return
	} else {
		fmt.Fprint(w, "Welcome to the dashboard")
	}

}
