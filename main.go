package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
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
	// 1. check the method
	if r.Method != "POST" {
		http.Error(w, "Bad Request", 400)
		return
	}

	// 2. read the two form values
	text := r.FormValue("text_input")
	banner := r.FormValue("banner_style")

	// 3. check for ascii validation
	splitText := strings.Split(text, "\\n")
	valid, err := validateInput(splitText)
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	// Load banner file
	bannerLine, err := loadBanner("banner/" + strings.ToLower(banner) + ".txt")
	if err != nil {
		http.Error(w, "Not Found", 404)
		return
	}

	// render ASCCI art
	generate := renderAll(valid, bannerLine)

	// send result back

}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	http.ListenAndServe(":8080", nil)
}
