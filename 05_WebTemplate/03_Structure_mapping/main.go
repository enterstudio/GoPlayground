package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: " This is a Title ",
		Body:  " This forms the Body ",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}
