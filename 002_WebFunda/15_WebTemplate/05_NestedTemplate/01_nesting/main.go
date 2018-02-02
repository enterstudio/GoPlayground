package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.html"))
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", nil)
}
