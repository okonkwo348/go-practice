package main

import (
	"fmt"
	"net/http"
)

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
		fmt.Fprint(w, "Welcome to version 2")
		return
	}
	if r.URL.Path == "/v2" {
		fmt.Fprint(w, "Welcome to version 2")
		return
	}

}
