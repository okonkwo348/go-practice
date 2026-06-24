package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")

}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if len(name) == 0 {
		fmt.Fprintln(w, "Greetings, Stranger!")
		return
	}

	fmt.Fprintf(w, "Greetings, %v!", name)

}
