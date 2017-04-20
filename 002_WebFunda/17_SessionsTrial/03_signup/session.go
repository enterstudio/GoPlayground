package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	c, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		//http.SetCookie(w, c) // Setting the Cookie
	}

	var u user
	if uid, ok := dbSession[c.Value]; ok {
		u = dbUsers[uid]
	}
	return u
}

func alreadyLogedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	e := dbSession[c.Value]
	_, ok := dbUsers[e] // Does the user really exists
	return ok
}
