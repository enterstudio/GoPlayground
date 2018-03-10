package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.ContentLength > 0 {
		bs := make([]byte, req.ContentLength)
		req.Body.Read(bs)
		s = string(bs)
	}
	tmpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 2", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, s)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error 3", http.StatusInternalServerError)
	}
}
