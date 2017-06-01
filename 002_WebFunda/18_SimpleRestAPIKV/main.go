package main

import (
	"fmt"
	"github.com/boseji/GoPlayground/002_WebFunda/18_SimpleRestAPIKV/handlers"
	"github.com/boseji/GoPlayground/002_WebFunda/18_SimpleRestAPIKV/storage"
	"google.golang.org/appengine"
	"net/http"
)

var Local bool

func init() {
	Local = appengine.IsDevAppServer()

	attachHandler(http.DefaultServeMux)
}

func main() {
	appengine.Main()
}

func attachHandler(h *http.ServeMux) {
	db := storage.NewInMemoryDB()
	h.HandleFunc("/", root)
	h.Handle("/favicon.ico", http.NotFoundHandler())
	h.HandleFunc("/get", handlers.GetKey(db))
	h.HandleFunc("/set", handlers.SetKey(db))
	h.HandleFunc("/del", handlers.DelKey(db))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("refresh", "5") // Refresh every 5 seconds
	fmt.Fprint(w, "Aum Sri Ganeshay Namh \nThis is an API Server")
	fmt.Fprint(w, "\n - Use GET Request '\\get?key=<Specific Key>' to access the Value stored")
	fmt.Fprint(w, "\n - Use PUT Request '\\set?key=<New Key> and in body <New Value>' to store the New Key value pair")
	fmt.Fprint(w, "\n - Use POST Request '\\set' and in body 'key = <New Key>''value = <New Value>' ")
	fmt.Fprint(w, "\n     encoded as 'application/x-www-form-urlencoded' to store the New Key value pair")
	fmt.Fprint(w, "\n - Use DELETE Request '\\del?key=<Specific Key>' to Delete the Key,Value pair")
	fmt.Fprint(w, "\n - In case the Key,Value pair previously exists they would be overwritten")
	fmt.Fprint(w, "\n - In case the Key,Value pair does not exist a Notfound message would be issued")
}
