package main

import (
	"net/http"
	"html/template"
	"log"
	"fmt"

)
var tmpl *template.Template
var err error

func homeHandler(w http.ResponseWriter, r *http.Request){
	tmpl.Execute(w, nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request){
	value:=r.FormValue("input_text")

	w.Write([]byte(value))
}


func main(){
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("http://localhost:8080")
	http.HandleFunc("GET /{$}", homeHandler)
	http.HandleFunc("POST /echo", echoHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}