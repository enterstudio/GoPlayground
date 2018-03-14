package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	cookieName  = "Site-Cookie1"
	cookieName1 = "NewAlternate"
	cookieName2 = "Special"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/read", read)
	http.HandleFunc("/multiple", multiple)
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

func multiple(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName1,
		Value: "This is Alternative Cookie",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName2,
		Value: "This is Special Cookie",
	})
	fmt.Fprint(w, "Multiple Cookie Written")
}

func read(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "Cookie #1 Value: %s", coo.Value)
	}
	coo, err = r.Cookie(cookieName1)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "\nCookie #2 Value: %s", coo.Value)
	}
	coo, err = r.Cookie(cookieName2)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintf(w, "\nCookie #3 Value: %s", coo.Value)
	}
}
