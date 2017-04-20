package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"html/template"
)

type user struct {
	UserName string
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
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, c) // Setting the Cookie
	}

	var u user
	if uid, ok := dbSession[c.Value]; ok {
		u = dbUsers[uid]
	}

	// Process the form
	if r.Method == http.MethodPost {
		e := r.FormValue("username")
		f := r.FormValue("first")
		l := r.FormValue("last")
		u = user{e, f, l}
		dbSession[c.Value] = e
		dbUsers[e] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	e, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[e]

	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
