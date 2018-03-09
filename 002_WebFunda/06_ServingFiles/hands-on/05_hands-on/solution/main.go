package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.HandleFunc("/", handler)
	http.Handle("/pics/", fs)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Println(err)
		http.Error(w, "File not found", 404)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
