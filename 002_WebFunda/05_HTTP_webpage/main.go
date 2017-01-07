package main

import (
	"io"
	"log"
	"net/http"
)

type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Aum Sri Ganeshay Namh !")
}

func main() {

	var h MyHandler
	log.Println("Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", h))
}
