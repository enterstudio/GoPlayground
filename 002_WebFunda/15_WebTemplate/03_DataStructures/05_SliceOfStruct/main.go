package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

type passingData struct {
	Name  string
	Quote string
}

func main() {
	p1 := passingData{
		Name:  "Krishna",
		Quote: "Submit on to me",
	}
	p2 := passingData{
		Name:  "Swami Vivekananda",
		Quote: "Arise Awake , stop not till the Goal is reached",
	}
	p3 := passingData{
		Name:  "Lala Lajpath Rai",
		Quote: "Swaraj is my Birthright",
	}
	pSlice := []passingData{
		p1,
		p2,
		p3,
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", pSlice)
}
