package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)

type Entity struct {
	Value string
}

func NumTables(ctx context.Context) (int, error) {
	s, err := datastore.Kinds(ctx)
	return len(s), err
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx := appengine.NewContext(r)

	// Find if there are any Tables available
	l, err := NumTables(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, " Database Available: %d", l)
}

func init() {
	http.HandleFunc("/", handle)
}
