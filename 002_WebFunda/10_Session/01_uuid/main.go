package main

import (
	"io"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	// Read Cookie
	cookie, err := req.Cookie("session-id")
	// If Cookie is not set
	if err != nil {
		id, _ := uuid.NewV4() // Generate the UUID
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String() + " email=dev@email.com" + " JSON data",
			// Secure: true,
			HttpOnly: true, // No Javascript access
		}
		http.SetCookie(res, cookie)
	}
	log.Println(cookie)
	io.WriteString(res, "Aum Sri Ganesh !")
}

func main() {

	http.HandleFunc("/", serveRoot)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Ignore the Fav Ico handler
	log.Println(" Starting Server on Port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
