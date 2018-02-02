package main

import (
	"os"
	"text/template"
)

type item struct {
	Name string
	Cost float32
}

type meal struct {
	Meal  string
	Items []item
}

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
	Menu    []meal
}

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	resturants := []hotel{
		hotel{
			"Name 1", "Address 1", "City 1", "123-1", "Southern",
			[]meal{
				meal{"Breakfast", []item{{"Dosa", 20}, {"Idli", 20}}},
				meal{"Lunch", []item{{"Utta", 50}, {"Roti-Curry", 40}}},
				meal{"Dinner", []item{{"Vangi-Bhat", 40}, {"Ratri Utta", 60}}},
			},
		},
		hotel{
			"Name 2", "Address 2", "City 2", "123-2", "Northern",
			[]meal{
				meal{"Lunch", []item{{"Utta", 50}, {"Roti-Curry", 40}}},
			},
		},
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", resturants)
}
