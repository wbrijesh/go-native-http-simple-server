package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404: File Not Found", http.StatusNotFound)
		return
	} 

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm()
	err != nil {
		fmt.Fprintf(w, "Error parsing form %v", err)
		return
	}

	fmt.Fprintf(w, "POST Request Successful")

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "email = %s\n", email)
}

func main() {
	var fileServer = http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at localhost:8080 \n")
	if err := http.ListenAndServe(":8080", nil)
	err != nil {
		log.Fatal("Error starting server /n", err)
	}
}