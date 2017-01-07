package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
)

// Create the Session Storage
var store = sessions.NewCookieStore([]byte("This is a Secret Key used in Encryption of Cookies"))

func main() {
	http.HandleFunc("/", serveRoot)
	log.Println("Starting Server on port 9000")
	// Special needed for Gorilla Sessions
	log.Fatalln(http.ListenAndServe(":9000", context.ClearHandler(http.DefaultServeMux)))
}

func serveRoot(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	if req.FormValue("email") != "" {
		session.Values["email"] = req.FormValue("email")
	}
	session.Save(req, res)

	htm := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
<h1> LOGIN: </h1>
<form method="post">
	<h3>` + fmt.Sprint(session.Values["email"]) + `</h3>
	<h3>Enter your Email:</h3>
	<input type="email", name="email">
	<input type="submit">
</form>
</body>
</html>
	`
	io.WriteString(res, htm)
}
