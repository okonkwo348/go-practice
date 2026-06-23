package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func statusHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", 400)
		return
	}

	codeNun, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", 400)
		return
	}

	if codeNun < 100 || codeNun > 599 {
		http.Error(w, "code must be a valid HTTP status code (100-599)", 400)
		return
	}

	if codeNun > 100 || codeNun < 599 {
		w.WriteHeader(codeNun)
		fmt.Fprintf(w, "Responding with status %d %s", codeNun, http.StatusText(codeNun))
		return
	}

}

/*Observation
when w.WriteHeader was placed before w.Write or fmt.Fprint the status was always 200 irrespective
of the code given at the query parameter */
