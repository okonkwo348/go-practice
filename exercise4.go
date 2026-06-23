package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	//What is the difference between r.ParseForm() + r.Form.Get()
	//  and just r.FormValue()?

	//The main difference is error control and access to different request parts.

	//r.FormValue() automatically call r.ParseForm() internally.However, if
	// the client sends a malformed body or a broken query string, r.FormValue()
	//ignoreds the error and returns an empty string "".

	/*r.ParseForm() + r.Form.Get(): Calling r.ParseForm() explicitly gives you total
	control over error handling. if a malicious or broken client sends corrupt data,
	you can catch it immediately and return a 400 Bad Request before your application
	processes anything else.*/

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}

	username := r.FormValue("username")
	language := r.FormValue("language")

	if len(username) == 0 {
		http.Error(w, "username is required", 400)
		return
	}

	if len(language) == 0 {
		http.Error(w, "language is required", 400)
		return
	}

	if len(username) != 0 && len(language) != 0 {
		fmt.Fprintf(w, "Hello %v, your are coding in %v!", username, language)
	}
}
