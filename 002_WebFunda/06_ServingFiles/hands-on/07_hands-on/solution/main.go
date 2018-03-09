package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", handler)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Println(err)
		http.Error(w, "File Not Found", 404)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
