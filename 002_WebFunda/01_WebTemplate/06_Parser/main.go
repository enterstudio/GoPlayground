package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title   string
	Heading string
	Content string
}

func main() {

	p1 := Page{
		Title:   "Go Test Page Parsing Example",
		Heading: " Go can be really Awesome ",
		Content: ` We can write Text or inject code
		<script>alert("Hello");</script>`,
	}

	tmp, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Println(err)
	}

	err = tmp.Execute(os.Stdout, p1)
	if err != nil {
		log.Println(err)
	}
}
