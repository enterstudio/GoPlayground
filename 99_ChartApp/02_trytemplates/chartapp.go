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

var templates = template.New("")

func init() {
	_, err := templates.ParseFiles("templates\\display.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	info1 := Page{Title: "Template Interpret", Body: "Testing Output"}
	fmt.Println(" First template t1:")
	err := templates.ExecuteTemplate(os.Stdout, "display.html", &info1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Total Number of Templates: ")
	for t := range templates.Templates() {
		fmt.Printf("\n - %d. %q", t, templates.Templates()[t].Name())
	}
	fmt.Println()
}
