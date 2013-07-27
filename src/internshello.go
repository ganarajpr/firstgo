package main

import (
	"fmt"
	"net/http"
	"log"
	//"strings"
	//"html/template"

)


func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	//:index.Execute(w, nil)
	if r.Method == "GET" {
		a := r.FormValue("username")
		fmt.Fprintf(w, "hi %s!",a); //<-- this variable does not rendered in the browser!!!

	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}


func main() {
	http.HandleFunc("/", GeneralHandler)
	http.HandleFunc("/hello/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil),nil)
}
