package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pagedata struct {
	Title    string
	Forms    bool
	Location string
	Method   string
	Data     string
}

var (
	pageroot = pagedata{"Home Page", false, "/", "", ""}
	pagepost = pagedata{"Post Form Page", true, "/process", "POST", ""}
	pageget  = pagedata{"Get Form Page", true, "/process", "GET", ""}
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/postform", poster)
	http.HandleFunc("/getform", getter)
	http.HandleFunc("/process", proc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Starting Server at 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, req *http.Request) {
	log.Printf("At root - Method Type: %s\n", req.Method)
	data := pageroot
	name := req.FormValue("na")
	prof := req.FormValue("prof")
	if len(name) != 0 && len(prof) != 0 {
		data.Data = fmt.Sprintf("Name = %s\nProfessions = %s", name, prof)
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Processing Error", http.StatusInternalServerError)
	}
}

func getter(w http.ResponseWriter, req *http.Request) {
	log.Printf("At /getform - Method Type: %s\n", req.Method)
	data := pageget
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Processing Error", http.StatusInternalServerError)
	}
}

func poster(w http.ResponseWriter, req *http.Request) {
	log.Printf("At /postform - Method Type: %s\n", req.Method)
	data := pagepost
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Processing Error", http.StatusInternalServerError)
	}
}

func proc(w http.ResponseWriter, req *http.Request) {
	log.Printf("At /process - Method Type: %s\n", req.Method)
	//name := req.FormValue("na")
	//prof := req.FormValue("prof")
	//log.Printf("Got Values \n %s,\n %s\n", name, prof)
	// Prepare a Get request URL
	//uri := fmt.Sprintf("/?na=%s&prof=%s", name, prof)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}
