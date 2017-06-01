package main

import (
	"fmt"
	"handlers"
	"net/http"
	"storage"
)

func init() {
	db := storage.NewInMemoryDB()
	http.HandleFunc("/", root)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/get", handlers.GetKey(db))
	http.HandleFunc("/set", handlers.SetKey(db))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("refresh", "5") // Refresh every 5 seocnds
	fmt.Fprint(w, "Aum Sri Ganeshay Namh \nThis is an API Server")
	fmt.Fprint(w, "\n - Use GET Request '\\get?key=<Specific Key>' to access the Value stored")
	fmt.Fprint(w, "\n - Use PUT Request '\\set?key=<New Key> and in body <New Value>' to store the New Key value pair")
	fmt.Fprint(w, "\n - In case the Key,Value pair previously exists they would be overwritten")
	fmt.Fprint(w, "\n - In case the Key,Value pair does not exist a Notfound message would be issued")
}
