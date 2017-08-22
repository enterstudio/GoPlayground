package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("01_text.gohtml")
	if err != nil {
		fmt.Println("Error in Finding template")
		return
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("Error in executing template")
	}
}
