package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request", 400)
		return
	}

	ts, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal Status Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Staus Error", http.StatusInternalServerError)
		return
	}
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHander)

	http.ListenAndServe(":8080", nil)
}
