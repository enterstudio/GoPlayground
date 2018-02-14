package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler1 int

func (h myHandler1) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "First Root Handler")
}

type myHandler2 int

func (h myHandler2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func main() {
	m := http.NewServeMux()
	var h1 myHandler1
	m.Handle("/", h1)
	var h2 myHandler2
	m.Handle("/about", h2)

	log.Println("Starting Server at 8080")
	log.Fatalln(http.ListenAndServe(":8080", m))
}
