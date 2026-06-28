package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.HandleFunc("/method-inspector", inspectorHandler)
	mainMux.HandleFunc("POST /echo", echoHandler)
	mainMux.HandleFunc("/headers", headersHandler)
	mainMux.HandleFunc("POST /form", formHandler)
	mainMux.HandleFunc("/status", statusHandler)
	mainMux.HandleFunc("/render", renderHandler)

	apiMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	fmt.Println("Serving on http://localhost:8080")

	err := http.ListenAndServe(":8080", mainMux)
	if err != nil {
		log.Fatal(err)
	}
}
