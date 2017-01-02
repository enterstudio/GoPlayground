package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	critical := color.New(color.BgHiRed, color.FgHiYellow)
	warning := color.New(color.BgHiYellow, color.FgHiBlue)
	info := color.New(color.FgHiBlue)
	highlight := color.New(color.FgHiWhite)
	success := color.New(color.BgHiGreen, color.FgHiBlue)

	critical.Println("\n This is a Critical Error Example ")
	warning.Println("\n This is a Serious Warning ")
	info.Println("\n This is an Info Message")
	highlight.Println("\n This is a Highlighted message")
	fmt.Println("\n This is a normal message ")
	success.Println("\n This is a Success message ")
}
