package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", name)
	log.Println("Starting Server on 8080 port")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, " We are at the Root Directory")
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, " We are at the Dog Directory")
}

func name(w http.ResponseWriter, req *http.Request) {
	pos := strings.LastIndex(req.URL.String(), "/me")
	//log.Printf("Last Position: %d", pos)

	if pos == -1 {
		log.Printf("Error Could not filter the URL String %s\n", req.URL.String())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	filtered := req.URL.String()[pos+4:]
	log.Printf("Filtered Text: %s", filtered)
	if len(filtered) == 0 {
		fmt.Fprintf(w, " Hello No Name !")
	} else {
		fmt.Fprintf(w, " Hello %s !", filtered)
	}
}
