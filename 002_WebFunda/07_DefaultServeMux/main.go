package main

import (
	"io"
	"log"
	"net/http"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Aum Sri Ganesh !")
}

func serveAum(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hare Hare")
}
func main() {

	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/aum/", serveAum)

	log.Println(" Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
