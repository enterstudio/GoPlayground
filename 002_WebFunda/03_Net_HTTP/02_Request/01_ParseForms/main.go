package main

import (
	"html/template"
	"log"
	"net/http"
)

type myHandler int

var tmpl *template.Template

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		tmpl.Execute(w, r.Form)
	} else {
		tmpl.Execute(w, nil)
	}
}

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h myHandler
	log.Println("Starting Server at port 8080")
	log.Fatalln(http.ListenAndServe(":8080", h))
}
