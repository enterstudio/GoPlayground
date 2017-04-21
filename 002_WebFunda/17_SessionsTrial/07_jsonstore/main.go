package main

import (
	"net/http"

	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"io/ioutil"
	"os"
	"os/signal"
)

type user struct {
	UserName string `json:"username"`
	Password []byte `json:"password"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

var tpl *template.Template

var dbUsers = map[string]user{}     // User ID, user
var dbSession = map[string]string{} // Session ID, User ID (email)

const (
	userFILE = "userStore.txt"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	// user data save and retrieval
	loadFromJson()
	defer saveToJson()

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	saveToJson()
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

		// Hash the password
		ps, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Store the User
		u := user{e, ps, f, l}
		dbUsers[e] = u
		fmt.Println("User Created ", u)

		// Done need to redirect to main page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, r *http.Request) {

	if alreadyLogedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Process the form
	if r.Method == http.MethodPost {
		e := r.FormValue("username")
		p := r.FormValue("password")

		// Is there a user?
		u, ok := dbUsers[e]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// Does the Entered password match that of the user
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
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

		fmt.Println("User Logged In: ", e)

		// Done need to redirect to main page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {

	if !alreadyLogedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := getUser(w, r)
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func logout(w http.ResponseWriter, r *http.Request) {

	if !alreadyLogedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")
	// Delete the Session
	delete(dbSession, c.Value)
	//Remove the Cookie data
	c = &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, c) // Setting the Cookie
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loadFromJson() {
	// Exit if the file does not exist
	if _, err := os.Stat(userFILE); err != nil {
		return
	}
	// If it does
	m, err := ioutil.ReadFile(userFILE)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(m, &dbUsers)
	if err != nil {
		panic(err)
	}
}

func saveToJson() {

	m, err := json.Marshal(dbUsers)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(userFILE, m, 0644)
	if err != nil {
		panic(err)
	}
}
