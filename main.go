package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Guarding the Gate (Request Verification)
	if r.Method != "GET" {
		http.Error(w, "Bad Request", 400)
		return
	}

	//Finding the HTML File (Parsing)
	ts, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal Status Error", http.StatusInternalServerError)
		return
	}

	// Delivering the Web Page (Executing)
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Staus Error", http.StatusInternalServerError)
		return
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POst" {
		http.Error(w, "Bad Request", 400)
		return
	}

}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	http.ListenAndServe(":8080", nil)
}
