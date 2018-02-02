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

	data := map[string]string{
		"Hari": `Aum`,
		"Hare": `Krishna`,
		"Ram":  `Sri Ram`,
		"Aum":  `Shanti`,
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println()
}
