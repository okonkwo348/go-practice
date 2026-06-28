package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", nameHandler)
	http.HandleFunc("/ping", homeHandler)
	http.HandleFunc(" /count", countHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/v2", legacyHandler)
	http.HandleFunc("/agent", agentHandler)

	fmt.Println("Server run on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
