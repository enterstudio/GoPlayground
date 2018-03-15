package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

type users struct {
	First string
	Last  string
	Email string
}

var tmpl *template.Template
var dbUser = map[string]users{}     // Storage of User details
var dbSession = map[string]string{} //Storage for Sessions created by users

func init() {
	tmpl = template.Must(template.ParseGlob("./templates/*.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", root)
	http.HandleFunc("/details", details)
	log.Println("Starting Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session")
	if err != nil {
		uid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		c = &http.Cookie{
			Name:  "session",
			Value: uid.String(),
		}
	}

	if r.Method == http.MethodPost {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		email := r.FormValue("mail")
		u := users{fname, lname, email}
		// Every time a new User is created a New Session shall be generated
		uid, err := uuid.NewV4()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		c.Value = uid.String()
		// Create or Update New user
		dbUser[email] = u
		dbSession[c.Value] = email

		log.Printf("Creating User with Session:\n %v\n %v", u, c.Value)
	}

	session := dbSession[c.Value]
	u, ok := dbUser[session]

	if ok {
		http.SetCookie(w, c)
	}
	err = tmpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		log.Println("Session Cookie does not Exists")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	session, ok := dbSession[c.Value]
	if !ok {
		log.Println("Session not Found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u, ok := dbUser[session]
	if !ok {
		log.Println("User Not found for Session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = tmpl.ExecuteTemplate(w, "details.gohtml", u)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
}
