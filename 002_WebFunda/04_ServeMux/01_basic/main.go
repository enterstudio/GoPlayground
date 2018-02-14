package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintln(w, "Home Page")
	case "/about":
		fmt.Fprintln(w, "About Page")
	default:
		w.WriteHeader(404)
	}
}

func main() {
	var h myHandler
	log.Println("Starting Server at 8080")
	log.Fatalln(http.ListenAndServe(":8080", h))
}
