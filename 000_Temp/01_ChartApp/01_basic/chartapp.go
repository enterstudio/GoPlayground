package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/hello", EchoHandler)
	http.HandleFunc("/hello/", EchoHandler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.Handle("/", http.FileServer(http.Dir("www")))
	log.Println("Starting Server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
