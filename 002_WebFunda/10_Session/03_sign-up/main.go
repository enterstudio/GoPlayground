package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	First    string
	Last     string
	UserName string
	Password string
}

var tmpl *template.Template
var dbUser = map[string]user{}      // Storage of User details
var dbSession = map[string]string{} //Storage for Sessions created by user

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/details", details)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	// Check if we Have the Session
	session, ok := dbSession[c.Value]
	if ok {
		// Check if the User Exists
		_, ok = dbUser[session]
		if ok {
			return true
		}
	}
	return false
}

func getUser(w http.ResponseWriter, r *http.Request) user {

	c, err := r.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
	}
	var u user
	if un, ok := dbSession[c.Value]; ok {
		u = dbUser[un]
	}
	return u
}

func root(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	err := tmpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("uname")
		pass := r.FormValue("pass")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")

		//If the Username Already exits
		if _, ok := dbUser[un]; ok {
			http.Error(w, "User name Already takn", http.StatusForbidden)
			return
		}

		// Add the User
		u = user{fname, lname, un, pass}
		dbUser[un] = u

		// create new session
		sid, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sid.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un

		// Redirect to Home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := tmpl.ExecuteTemplate(w, "signup.gohtml", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}

	err := tmpl.ExecuteTemplate(w, "details.gohtml", u)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}
