package main

import (
	"os"
	"text/template"
)

type item struct {
	Name string
	Cost float32
}

type menu struct {
	MenuName string
	Items    []item
}

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	resturantMenu := []menu{
		menu{"Breakfast", []item{{"Dosa", 20}, {"Idli", 20}}},
		menu{"Lunch", []item{{"Utta", 50}, {"Roti-Curry", 40}}},
		menu{"Dinner", []item{{"Vangi-Bhat", 40}, {"Ratri Utta", 60}}},
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", resturantMenu)
}
