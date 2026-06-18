package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	add := r.URL.Query().Get("add")
	substract := r.URL.Query().Get("subtract")
	multiply := r.URL.Query().Get("multiply")

	Anun, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	Bnun, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	if substract == "substract" {
		sub := Anun - Bnun
		fmt.Fprint(w, sub)
		return
	} else if add == "add" {
		add := Anun + Bnun
		fmt.Fprint(w, add)
		return
	} else if multiply == "multiply" {
		mul := Anun * Bnun
		fmt.Fprint(w, mul)
		return
	} else {
		http.Error(w, "20", http.StatusBadRequest)
		return
	}
}
