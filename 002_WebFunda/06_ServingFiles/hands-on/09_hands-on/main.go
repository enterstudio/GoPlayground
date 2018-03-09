package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("starting-files")))
	//http.Handle("/pic/", http.StripPrefix("/pic", http.FileServer(http.Dir("starting-files/pic"))))
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
