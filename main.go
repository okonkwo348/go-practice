package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type ArtResponse struct {
	UserInput   string
	AsciiArt    string
	BannerStyle string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Guarding the Gate (Request Verification)

	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Allow", 405)
		return
	}

	//Finding the HTML File (Parsing)
	ts, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		http.Error(w, "internal Status Error", http.StatusInternalServerError)
		return
	}

	// Delivering the Web Page (Executing)
	err = ts.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to execute template: %v", err)
		http.Error(w, "Internal Staus Error", http.StatusInternalServerError)
		return
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// 1. check the method
	if r.Method != "POST" {
		http.Error(w, "Method Not Allow", 405)
		return
	}

	// fmt.Println(r)

	// 2. read the two form values
	text := r.FormValue("text_input")
	banner := r.FormValue("banner_style")

	// check if text is empty
	if text == "" {
		http.Error(w, "Please Type at Least a Word AND select a banner", 400)
		return
	}

	// 3. check for ascii validation
	splitText := strings.Split(text, "\\n")
	valid, err := validateInput(splitText)
	if err != nil {
		http.Error(w, "Bad request: Does not accept non ascii characters", 400)
		return
	}

	// check if banner was selected
	if banner == "" {
		http.Error(w, "Select a banner", 400)
		return
	}

	// Load banner file
	bannerLine, err := loadBanner("banner/" + strings.ToLower(banner) + ".txt")
	if err != nil {
		http.Error(w, "Banner Not Found", 404)
		return
	}

	// render ASCCI art
	generate := renderAll(valid, bannerLine)

	responseData := ArtResponse{
		UserInput:   text,
		AsciiArt:    generate,
		BannerStyle: banner,
	}

	// send result back
	ts, err := template.ParseFiles("templates/result.html")
	if err != nil {
		log.Printf("Failed to parse template: %v", err)
		http.Error(w, "Internal Status Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, responseData)
	if err != nil {
		log.Printf("Failed to execute template: %v", err)
		http.Error(w, "Internal Status Error", http.StatusInternalServerError)
		return
	}

}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	log.Println("Server start on http://localhost:8080")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
