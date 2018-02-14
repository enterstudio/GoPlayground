package main

import (
	"io"
	"log"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Aum Sri Ganeshay Namh !")
}

func main() {
	var h myHandler
	log.Println("Starting Server on Port 8080")
	log.Fatalln(http.ListenAndServe(":8080", h))
}
