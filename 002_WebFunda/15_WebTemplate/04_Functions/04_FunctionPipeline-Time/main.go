package main

import (
	"os"
	"strings"
	"text/template"
	"time"
)

func timeformat(t time.Time) string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return t.In(loc).Format("02-01-2006 15:04:05 IST")
}

func addon(tm string) string {
	sp := strings.Split(tm, " ")
	sp[0] = "Date = " + sp[0]
	sp[1] = "time = " + sp[1]
	sp[2] = "TimeZone = " + sp[2]
	jn := strings.Join(sp[:3], " ")
	if len(sp) > 3 {
		jn += sp[3]
	}
	return jn
}

var fnpar = template.FuncMap{
	"fdate": timeformat,
	"addon": addon,
}

func main() {
	tpl := template.Must(template.New("").Funcs(fnpar).ParseFiles("tpl.gohtml"))
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		panic(err)
	}
}
