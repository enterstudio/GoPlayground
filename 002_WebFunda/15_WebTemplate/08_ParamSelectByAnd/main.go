package main

import (
	"os"
	"text/template"
)

var Details = []struct {
	Name    string
	Yuga    string
	Vanvasa bool
	Brother string
}{
	{"Ram", "Treta", true, "Lakshmana"},
	{"Lakshmana", "Treta", true, "Ram"},
	{"Krishna", "Dwaper", false, "Balarama"},
	{"Balarama", "Dwaper", false, "Krishna"},
}

func main() {
	tpl := template.Must(template.New("").ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Details)
	if err != nil {
		panic(err)
	}
}
