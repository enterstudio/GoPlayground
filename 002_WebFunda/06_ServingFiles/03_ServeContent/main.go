package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/puppy.jpg", puppy)
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

func puppy(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("images/puppy.jpg")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Special", "This is a JPEG Image interpret by browser")
	http.ServeContent(w, req, stat.Name(), stat.ModTime(), file)
	file.Close()
}
