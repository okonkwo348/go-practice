package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", 404)
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Page Not Found", 404)
		return
	}

	text := r.FormValue("input_text")
	banner := r.FormValue("banner_style")

	bannerFile := banner + ".txt"

	if len(text) == 0 || len(banner) == 0 {
		http.Error(w, "bad request", 400)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "bad request", 400)
		return
	}

	generate, err := GenerateArt(text, bannerFile)
	if err != nil {
		http.Error(w, "Internal Error server", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Result: generate,
		Text:   text,
		Banner: banner,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

}

func SwitcHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-switch" {
		http.Error(w, "Page Not Found", 404)
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if len(text) == 0 || len(banner) == 0 {
		http.Error(w, "bad request", 400)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "bad request", 400)
		return
	}

	bannerFile := banner + ".txt"
	generate, err := GenerateArt(text, bannerFile)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	fmt.Println(generate, text)

	data := PageData{
		Result: generate,
		Text:   text,
		Banner: banner,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeHandler)
	mux.HandleFunc("POST /ascii-art", AsciiHandler)
	mux.HandleFunc("GET /ascii-art-switch", SwitcHandler)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
