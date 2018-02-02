package main

import (
	"math"
	"os"
	"text/template"
)

var fnmap = template.FuncMap{
	"sqrt":   math.Sqrt,
	"square": square,
	"cube":   cube,
}

func square(n float64) float64 {
	return math.Pow(n, 2)
}

func cube(n float64) float64 {
	return math.Pow(n, 3)
}

func main() {
	tpl := template.Must(template.New("").Funcs(fnmap).ParseFiles("tpl.html"))
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", 2.0)
}
