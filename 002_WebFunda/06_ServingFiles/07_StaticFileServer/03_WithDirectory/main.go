package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Starting File Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
