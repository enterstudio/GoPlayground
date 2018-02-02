package main

import (
	"os"
	"text/template"
	"time"
)

var fnpar = template.FuncMap{
	"fdate": timeformat,
}

func timeformat(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return t.In(loc).Format("02-January-2006 15:04:05 IST")
}

func main() {
	tpl := template.Must(template.New("").Funcs(fnpar).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now().Add(24*time.Hour))
	if err != nil {
		panic(err)
	}
}
