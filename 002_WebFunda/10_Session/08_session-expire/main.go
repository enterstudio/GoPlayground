package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
	Password []byte
	Role     string
}

type session struct {
	un          string
	lastActivty time.Time
}

const (
	cookieName   = "session"
	adminRole    = "admin"
	expiryPeriod = 30 // Seconds
	expiryTime   = 30 * time.Second
)

var tmpl *template.Template
var dbUsers = map[string]user{}       // Storage for User data
var dbSessions = map[string]session{} // Storage for Session data
var lastCleanup = time.Now()

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
		sess := dbSessions[c.Value]
		_, ok := dbUsers[sess.un]
		if ok {
			return true
		}
	}

	return false
}

func getUserFromSession(w http.ResponseWriter, r *http.Request) user {
	var u user
	c, err := r.Cookie(cookieName)
	if err == nil {
		sess, ok := dbSessions[c.Value]
		u, ok = dbUsers[sess.un]
		// If the User is geniuine
		if ok {
			// Update timeout
			c.MaxAge = expiryPeriod
			http.SetCookie(w, c)
			sess.lastActivty = time.Now()
		}
	}
	return u
}

func createSession(w http.ResponseWriter, r *http.Request, u user) {
	sID, _ := uuid.NewV4()
	c := &http.Cookie{
		Name:   cookieName,
		Value:  sID.String(),
		MaxAge: expiryPeriod,
	}
	http.SetCookie(w, c)
	sess := session{u.UserName, time.Now()}
	dbSessions[c.Value] = sess

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func cleanupSession() {
	// Expired
	if lastCleanup.Add(expiryTime).Before(time.Now()) {
		log.Println("************** BEFORE")
		log.Println(dbSessions)

		for k, v := range dbSessions {
			// Session Expired
			if v.lastActivty.Add(expiryTime).Before(time.Now()) {
				delete(dbSessions, k)
			}
		}

		log.Println("************** AFTER")
		log.Println(dbSessions)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	u := getUserFromSession(w, r)
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

		// Invalid Password - bcrypt match
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(pass))
		if err != nil {
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

		// Create new ID & login the User - bcrypt password
		bpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		u := user{un, fname, lname, bpass, role}
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
	u := getUserFromSession(w, r)
	tmpl.ExecuteTemplate(w, "details.gohtml", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadySignedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie(cookieName)
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	// Run Cleanup at logout
	cleanupSession()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func special(w http.ResponseWriter, r *http.Request) {
	if !alreadySignedIn(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	u := getUserFromSession(w, r)
	if u.Role != adminRole {
		http.Error(w, "You need to have special permissions", http.StatusForbidden)
		return
	}
	tmpl.ExecuteTemplate(w, "special.gohtml", u)
}
