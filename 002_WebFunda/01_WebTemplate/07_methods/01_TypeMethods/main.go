package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

// Note: No Pointer Receivers allowed for Template Functions
//       Protection measure for not Manipulating data by mistake

func (p person) SomeProcess() int {
	return 7
}

func (p person) AgeDouble() int {
	return 2 * p.Age
}

func (p person) ToString(age int) string {
	return fmt.Sprintf(" %s is of Age %d", p.Name, age)
}

func main() {
	p := person{"Radha", 20}
	tpl := template.Must(template.ParseGlob("*.html"))
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", p)
}
