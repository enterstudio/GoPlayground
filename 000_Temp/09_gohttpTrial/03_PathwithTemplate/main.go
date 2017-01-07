package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("template/index.gohtml")
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}
