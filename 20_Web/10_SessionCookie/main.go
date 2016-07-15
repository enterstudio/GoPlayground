package main

import (
	"io"
	"log"
	"net/http"
	"github.com/nu7hatch/gouuid"
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
		}
		http.SetCookie(res, cookie)
	}
	log.Println(cookie)
	io.WriteString(res, "Aum Sri Ganesh !")
}

func main() {

	http.HandleFunc("/", serveRoot)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Ignore the Fav Ico handler
	log.Println(" Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
