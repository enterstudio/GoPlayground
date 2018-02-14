package main

import (
	"io"
	"log"
	"net/http"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Aum Sri Ganeshay Namh !")
}

func serveAum(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hare Hare")
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", serveRoot)
	mux.HandleFunc("/aum/", serveAum)

	log.Println(" Starting Server on Port 8080")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
