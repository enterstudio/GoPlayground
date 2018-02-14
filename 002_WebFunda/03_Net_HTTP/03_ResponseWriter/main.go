package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("NewKey", "New Value for the New Key")
	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	fmt.Fprintln(w, "<h1>This is Content generated </h1>")
}

func main() {
	var h myHandler

	log.Println("Starting Server at 8080")
	log.Fatalln(http.ListenAndServe(":8080", h))
}
