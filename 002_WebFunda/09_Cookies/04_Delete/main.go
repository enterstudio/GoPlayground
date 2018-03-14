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
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `<h1><a href="/set">Set</a></h1>`)
}

func set(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: "This is a Site Cookie created",
	})
	fmt.Fprint(w, `<h1><a href="/read">Read</a></h1>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, "<h1>Cookie Value: %s</h1>", coo.Value)
	fmt.Fprint(w, `<h1><a href="/expire">Expire</a></h1>`)
}

func expire(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	coo.MaxAge = -1 // Delete the Cookie
	http.SetCookie(w, coo)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
