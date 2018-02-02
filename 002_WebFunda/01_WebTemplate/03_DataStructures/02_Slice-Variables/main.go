package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", []int{1, 3, 5})
	fmt.Print("\n\n\n Now using a Variable\n\n")
	v := []int{20, 30, 40}
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", v)
}
