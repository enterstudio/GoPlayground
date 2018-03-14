package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const cookieName = "Site-Cookie1"

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/read", read)
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	visit := 0
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		cookie = &http.Cookie{
			Name:  cookieName,
			Value: "1",
		}
		http.SetCookie(w, cookie)
		fmt.Fprint(w, "Cookie Written\n")
	} else {
		visit, err = strconv.Atoi(cookie.Value)
		if err != nil {
			log.Fatalln(err)
		}
		visit++
		cookie.Value = fmt.Sprintf("%d", visit)
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Visited %d times", visit)
}

func read(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie(cookieName)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Cookie Value: %s", coo.Value)
}
