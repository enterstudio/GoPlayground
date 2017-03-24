package appengine_test

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Aum Sri Ganeshay Nmah!")
}
