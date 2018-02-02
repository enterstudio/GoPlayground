package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	hotels := []hotel{
		hotel{"Name 1", "Address 1", "City 1", "123-1", "Southern"},
		hotel{"Name 2", "Address 2", "City 2", "123-2", "Northern"},
		hotel{"Name 3", "Address 3", "City 3", "123-3", "Central"},
		hotel{"Name 4", "Address 4", "City 4", "123-4", "Southern"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}

}
