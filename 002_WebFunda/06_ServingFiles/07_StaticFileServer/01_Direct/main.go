package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting File Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
