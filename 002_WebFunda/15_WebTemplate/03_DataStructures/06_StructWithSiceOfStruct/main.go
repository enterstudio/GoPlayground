package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type FreedomFigher struct {
	Name  string
	Quote string
}

type State struct {
	Name        string
	CapitalCity string
}

type DataPassed struct {
	FreedomFighters []FreedomFigher
	States          []State
}

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	f1 := FreedomFigher{
		Name:  "Subash Chandra Bose",
		Quote: "Give me blood, I grant you Freedom",
	}
	f2 := FreedomFigher{
		Name:  "Lala Lajpath Rai",
		Quote: "Swaraj is my Birthright",
	}
	s1 := State{
		Name:        "Maharashtra",
		CapitalCity: "Mumbai",
	}
	s2 := State{
		Name:        "Tamil Nadu",
		CapitalCity: "Chennai",
	}
	d := DataPassed{
		FreedomFighters: []FreedomFigher{
			f1,
			f2,
		},
		States: []State{
			s1,
			s2,
		},
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", d)
}
