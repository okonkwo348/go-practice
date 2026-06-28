package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	operation:=r.URL.Query().Get("op")

	Anun, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	Bnun, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	switch operation {
	case r.URL.Query().Get("add"):
		sub := Anun - Bnun
		fmt.Fprint(w, sub)
		return
	case r.URL.Query().Get("subtract"):
		add := Anun + Bnun
		fmt.Fprint(w, add)
		return
	case r.URL.Query().Get("multiply"):
		mul := Anun * Bnun
		fmt.Fprint(w, mul)
		return
	default:
		 http.Error(w, "20", http.StatusBadRequest)
		return
	}
}
