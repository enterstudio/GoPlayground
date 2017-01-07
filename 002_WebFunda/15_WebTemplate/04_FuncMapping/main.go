package main

import (
	"os"
	"text/template"
	"time"
)

func main() {
	tpl := template.Must(template.New("").Funcs(fnpar).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		panic(err)
	}
}

func timeformat(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return t.In(loc).Format("02-01-2006 15:04:05 IST")
}

var fnpar = template.FuncMap{
	"fdate": timeformat,
}
