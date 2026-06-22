package main

import (
	"fmt"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	head := r.Header.Get("X-Custom-Token")
	if len(head) == 0 {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}

	if r.Header.Get("X-Custom-Token") == "abc123" {
		if len(r.Header.Get("Content-Type")) == 0 {
			fmt.Fprintf(w, "Token received: abc123\n%v", r.Header.Get("Content-Type"))

		} else {
			fmt.Fprint(w, "Token received: abc123\nContent-Type: application/json")
		}

	}
}
