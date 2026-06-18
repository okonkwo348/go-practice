package main

import (
	"fmt"
	"net/http"
)

func agentHandler(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	fmt.Fprint(w, agent)

}
