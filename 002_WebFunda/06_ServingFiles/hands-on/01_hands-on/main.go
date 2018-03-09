package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/puppy.jpg", filehandle)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "foo ran")
}

func dog(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("dog.gohtml")
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

func filehandle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "puppy.jpg")
}
