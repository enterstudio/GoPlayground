package main

import (
	"io"
	"log"
	"net/http"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	// Check If previous cookie in Response
	if cookie, err := req.Cookie("self-cookie"); err == nil {
		log.Println("Previous Cookie: ", cookie)
	} else {
		http.SetCookie(res, &http.Cookie{
			Name:  "self-cookie",
			Value: "This is my First Cookie",
		})
		log.Println("Setting up new Cookie: self-cookie")
	}
	io.WriteString(res, "Aum Sri Ganesh !")
}

func main() {

	http.HandleFunc("/", serveRoot)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Ignore the Fav Ico handler
	log.Println(" Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
