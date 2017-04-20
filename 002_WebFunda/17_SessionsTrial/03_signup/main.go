package main

import (
	"net/http"

	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template

var dbUsers = map[string]user{}     // User ID, user
var dbSession = map[string]string{} // Session ID, User ID (email)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func signup(w http.ResponseWriter, r *http.Request) {

	if alreadyLogedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Process the form
	if r.Method == http.MethodPost {
		e := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("first")
		l := r.FormValue("last")

		// Is user name already taken, user has already signed up ?
		if _, ok := dbUsers[e]; ok {
			http.Error(w, "Username Already taken", http.StatusForbidden)
			return
		}

		// Create new Session
		id := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c) // Setting the Cookie

		dbSession[c.Value] = e

		// Store the User
		u := user{e, p, f, l}
		dbUsers[e] = u
		fmt.Println("User Created ", u)

		// Done need to redirect to main page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {

	if !alreadyLogedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
