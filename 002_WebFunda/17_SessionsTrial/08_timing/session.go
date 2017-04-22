package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
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
		http.SetCookie(w, c) // Setting the Cookie
	}

	var u user
	if uid, ok := dbSession[c.Value]; ok {
		u = dbUsers[uid.UserId]
	}
	return u
}

func alreadyLogedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	mutSession.Lock()
	defer mutSession.Unlock()
	e, ok := dbSession[c.Value]
	if !ok {
		return false
	}
	tnow := time.Now().Unix() // Current Unix Time
	_, ok = dbUsers[e.UserId] // Does the user really exists
	if tnow > (int64(e.ValidTime) + e.StartTime) {
		delete(dbSession, c.Value)
		fmt.Println("User Logged Out(Timeout): ", e.UserId)
		ok = false // Since we now have the session Cleared
	} else {
		// Time has not happened yet on the session
		// Update the time stamp
		e.StartTime = tnow
		dbSession[c.Value] = e // Update the Session
	}
	return ok
}
