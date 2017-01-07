package main

import (
	"io"
	"log"
	"net/http"
)

func serveRoot(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/res/gopher.jpg">`)
}

func main() {

	http.HandleFunc("/", serveRoot)
	http.Handle("/res/", http.StripPrefix("/res", http.FileServer(http.Dir("./img"))))

	log.Println(" Starting Server on Port 9000")
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
