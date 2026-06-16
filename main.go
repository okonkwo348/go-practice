package main

import (
	"net/http"
	//"log"
	"fmt"
)
// func homeHandler(w http.ResponseWriter, r *http.Request){
// 	switch r.URL.Path{
// 	case "/ok":
// 		http.Error(w,"sucessful", 200)
// 	
// 	case "/badrequest":
// 		http.Error(w,"Bad Request", 400)
// 	
// 	case "/error":
// 		http.Error(w,"Internal Error", 500)
// 	
// 	default:
// 		http.Error(w,"Not Found", 404)
// 	
// 	}
// }

// func okHandler(w http.ResponseWriter, r *http.Request){
// 		http.Error(w, "Successful", 200)
		
// }

// func nfHandler(w http.ResponseWriter, r *http.Request){
// 		http.Error(w, "Not Found", 404)
		
// }

// func brHandler(w http.ResponseWriter, r *http.Request){
// 		http.Error(w, "Bad Request", 400)
		
// }

// func erHandler(w http.ResponseWriter, r *http.Request){
// 	http.Error(w, "500 internal server error", 500)
	
	
// }

// func main(){
// 	http.HandleFunc("/ok", okHandler)
// 	http.HandleFunc("/notfound", nfHandler)
// 	http.HandleFunc("/badrequest", brHandler)
// 	http.HandleFunc("/error", erHandler)
// 	fmt.Println("Server running on http//:localhost:8080")
// 	http.ListenAndServe(":8080",nil)
// }

func main(){
	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Successful", 200)
	})
	http.HandleFunc("/notfound", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Not Found", 404)
	})
	http.HandleFunc("/badrequest", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "Bad Request", 400)
	})
	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request){
		http.Error(w, "500 internal server error", 500)
	})
	fmt.Println("Server running on http//:localhost:8080")
	http.ListenAndServe(":8080",nil)
}