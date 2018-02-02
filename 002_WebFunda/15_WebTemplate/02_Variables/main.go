package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", 108)
}
