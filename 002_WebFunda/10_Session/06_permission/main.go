package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
	Password string
	Role     string
}

const (
	cookieName = "session"
	adminRole  = "admin"
)

var tmpl *template.Template
var dbUsers = map[string]user{}      // Storage for User data
var dbSessions = map[string]string{} // Storage for Session data

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/details", details)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/special", special)
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func alreadySignedIn(r *http.Request) bool {
	c, err := r.Cookie(cookieName)
	if err == nil {
		un := dbSessions[c.Value]
		_, ok := dbUsers[un]
		if ok {
			return true
		}
	}

	return false
}

func getUserFromSession(r *http.Request) user {
	var u user
	c, err := r.Cookie(cookieName)
	if err == nil {
		un := dbSessions[c.Value]
		u = dbUsers[un]
	}
	return u
}

func createSession(w http.ResponseWriter, r *http.Request, u user) {
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:  cookieName,
		Value: sID.String(),
	}
	http.SetCookie(w, c)
	dbSessions[c.Value] = u.UserName
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func root(w http.ResponseWriter, r *http.Request) {
	u := getUserFromSession(r)
	tmpl.ExecuteTemplate(w, "index.gohtml", u)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadySignedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("uname")
		pass := r.FormValue("pass")

		// Check if the username Actually Exists
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Wrong Inputs", http.StatusForbidden)
			return
		}

		// Invalid Password
		if u.Password != pass {
			http.Error(w, "Wrong Inputs", http.StatusForbidden)
			return
		}

		createSession(w, r, u)
		return
	}
	tmpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadySignedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		un := r.FormValue("uname")
		pass := r.FormValue("pass")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		role := r.FormValue("role")

		// Check if the username is previously taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username Already taken", http.StatusForbidden)
			return
		}

		// Create new ID & login the User
		u := user{un, fname, lname, pass, role}
		dbUsers[un] = u
		createSession(w, r, u)
		return
	}
	tmpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func details(w http.ResponseWriter, r *http.Request) {
	if !alreadySignedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUserFromSession(r)
	tmpl.ExecuteTemplate(w, "details.gohtml", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadySignedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUserFromSession(r)
	delete(dbSessions, u.UserName)
	c := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func special(w http.ResponseWriter, r *http.Request) {
	if !alreadySignedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUserFromSession(r)
	if u.Role != adminRole {
		http.Error(w, "You need to have special permissions", http.StatusForbidden)
		return
	}
	tmpl.ExecuteTemplate(w, "special.gohtml", u)
}
