package main

import (
	"os"
	"text/template"
)

var name = []string{"Krishna", "Ram", "Balaram", "Madhav"}

func main() {
	tpl := template.Must(template.New("").ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", name)
	if err != nil {
		panic(err)
	}
}
