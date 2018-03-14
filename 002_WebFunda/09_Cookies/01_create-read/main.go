package main

import (
	"fmt"
	"log"
	"net/http"
)

const cookieName = "Site-Cookie1"

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/read", read)
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: "This is a Site Cookie created",
	})
	fmt.Fprint(w, "\nCookie Written")
}

func read(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Cookie Value: %s", coo.Value)
}
