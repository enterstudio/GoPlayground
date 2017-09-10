package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	server := os.Getenv("C9_HOSTNAME")
	if len(server) == 0 {
		server = "0.0.0.0"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("template/index.gohtml")
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/pics/", fs)

	log.Println("Running Server on " + server + ":" + port)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
