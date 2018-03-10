package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler)
	log.Println("Starting server on 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	var bs []byte

	log.Println(req.Method)
	if req.Method == http.MethodPost {
		fl, header, err := req.FormFile("q")
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error1", http.StatusInternalServerError)
			return
		}

		defer fl.Close()
		bs, err = ioutil.ReadAll(fl)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error2", http.StatusInternalServerError)
			return
		}

		log.Println(header)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<form method="POST" enctype="multipart/form-data">
		Set File Name: <input type="file" name="q"><br>
		<input type="submit">
		</form>
		<p>%s</p>
		`, string(bs))
}
