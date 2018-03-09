package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
