package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle1)
	http.HandleFunc("/new", handle2)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Started server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handle1(w http.ResponseWriter, req *http.Request) {
	log.Printf("At Root - Method %s", req.Method)
}

func handle2(w http.ResponseWriter, req *http.Request) {
	log.Printf("At /new - Method %s", req.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
	// Or Shorter
	//http.Redirect(w, req, "/", http.StatusMovedPermanently)
}
