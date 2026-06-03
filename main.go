package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Wrong Method: needs GET method")
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Print("Wrong Method: POST Method allowed only")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text_input")
	banner := r.FormValue("banner_style")

	if text == "" {
		log.Print("Empty Banner")
		http.Error(w, "Input a string and select a banner.", 400)
		return
	}

	wordSplit := strings.Split(text, "\\n")
	valid, err := ValidateInput(wordSplit)
	if err != nil {
		log.Print("Character Not Found in the Ascii table")
		http.Error(w, "Character Not Found in the Ascii table", 400)
		return
	}

	if banner == "" {
		log.Print("Empty banner")
		http.Error(w, "Select a banner", 400)
		return
	}

	bannerfile, err := LoadBanner("banner/" + strings.ToLower(banner) + ".txt")
	if err != nil {
		log.Print("File Not Found")
		http.Error(w, "File Not Found", 404)
		return
	}

	data := RenderAll(valid, bannerfile)

	ts, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)

	log.Println("Server start on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
}
