package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, []int{1, 2, 3, 4})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}
