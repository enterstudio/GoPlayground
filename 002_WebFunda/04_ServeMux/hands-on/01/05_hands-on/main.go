package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tmpl *template.Template

type pagedata struct {
	Title   string
	Content string
}

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	//http.HandleFunc("/", home)
	http.Handle("/", http.HandlerFunc(home))
	//http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	//http.HandleFunc("/me/", name)
	http.HandleFunc("/me/", http.HandlerFunc(name))
	log.Println("Starting Server on 8080 port")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	data := pagedata{
		Title:   "Home Page",
		Content: "This is the Home Page",
	}
	err := tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}

func dog(w http.ResponseWriter, req *http.Request) {
	data := pagedata{
		Title:   "Dog Page",
		Content: "This is the Dog Page",
	}
	err := tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}

func name(w http.ResponseWriter, req *http.Request) {
	pos := strings.LastIndex(req.URL.String(), "/me")
	//log.Printf("Last Position: %d", pos)

	if pos == -1 {
		log.Printf("Error Could not filter the URL String %s\n", req.URL.String())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	filtered := req.URL.String()[pos+4:]
	log.Printf("Filtered Text: %s", filtered)
	if len(filtered) == 0 {
		filtered = "No Name"
	}

	data := pagedata{
		Title:   "Name Page",
		Content: fmt.Sprintf("Hello %s !", filtered),
	}
	err := tmpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Println(err)
	}
}
