package main

import (
	"fmt"
	"io"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		l := len(string(body))
		fmt.Fprint(w, l)
		return
	}

}
