package main

import (
	"io"
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
	v := req.FormValue("q")
	io.WriteString(w, "Here is Query Value q = "+v)
}
