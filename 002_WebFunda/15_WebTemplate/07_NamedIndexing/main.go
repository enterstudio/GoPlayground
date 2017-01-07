package main

import (
	"os"
	"text/template"
)

var name = []string{"Krishna", "Ram", "Balaram", "Madhav"}

var all = struct {
	Names   []string
	Message string
}{
	name,
	"Hare Krishna !",
}

func main() {
	tpl := template.Must(template.New("").ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", all)
	if err != nil {
		panic(err)
	}
}
