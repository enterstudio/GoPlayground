package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	log.Println("Starting File Server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<img src="/images/puppy.jpg">`)
}
