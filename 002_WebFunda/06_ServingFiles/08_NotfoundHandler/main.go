package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprint(w, "Start Looking at your Terminal")
}
