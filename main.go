package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/method-inspector", inspectorHandler)
	mux.HandleFunc("POST /echo", echoHandler)
	mux.HandleFunc("/headers", headersHandler)
	mux.HandleFunc("POST /form", formHandler)
	mux.HandleFunc("/status", statusHandler)
	fmt.Println("Serving on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
