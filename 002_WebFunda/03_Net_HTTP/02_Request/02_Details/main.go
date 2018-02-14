package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type myHandler int

var tmpl *template.Template

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.ContentLength,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
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
