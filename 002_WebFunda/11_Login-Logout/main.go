package main

import (
	"github.com/nu7hatch/gouuid"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", authServe)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // Ignore the Fav Ico handler
	log.Println(" Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}

func authServe(res http.ResponseWriter, req *http.Request) {

	// Get Session Cookie
	ses_cookie, err := req.Cookie("session-id")
	// If Cookie is not set
	if err != nil {
		id, _ := uuid.NewV4() // Generate the UUID
		ses_cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
		}
		http.SetCookie(res, ses_cookie)
	}
	log.Println(ses_cookie)

	// Get Login Cookie
	lgn_cookie, err := req.Cookie("login")
	// If cookie was not found
	if err != nil {
		lgn_cookie = &http.Cookie{
			Name:  "login",
			Value: "0",
		}
	}

	// Check If the Form was submitted for Login
	if req.Method == "POST" && req.URL.Path == "/" {
		pass := req.FormValue("password")
		// Check for Correct Password
		if pass == "secret" {
			lgn_cookie = &http.Cookie{
				Name:  "login",
				Value: "1",
			}
		}
	}

	// Check If Logout was clicked
	if req.URL.Path == "/logout" {
		lgn_cookie = &http.Cookie{
			Name:   "login",
			Value:  "0",
			MaxAge: -1,
		}
		http.SetCookie(res, lgn_cookie)
		// Setup a Redirect rather than the Page
		http.Redirect(res, req, "/", 302)
		log.Println(" Redirecting after Logout")
		return // Exit no need as the final web page would be re-generated
	}

	http.SetCookie(res, lgn_cookie)
	log.Println(lgn_cookie)

	// Prepare Content
	var htm string
	if lgn_cookie.Value == "0" { // In case we are Not Logged In
		htm = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
<h1> LOGIN: </h1>
<form method="post">
	<h3>User Name:</h3>
	<input type="text" name="userName">
	<h3>Password:</h3>
	<input type="text", name="password">
	<br>
	<input type="submit">
</form>
</body>
</html>
		`
	} else if lgn_cookie.Value == "1" { // In case we are Logged In
		htm = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	<h1> Welcome </h1>
	<a href="/logout"><h3> LOGOUT </h3></a>
</body>
</html>
		`
	}

	// Page Content
	io.WriteString(res, htm)
}
