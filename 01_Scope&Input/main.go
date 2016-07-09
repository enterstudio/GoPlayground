package main

import (
	"./submod"
	"fmt"
)

// checkker := true // This is wrong as the ':=' can only be used in functions


func main() {
	var name string
	fmt.Printf("\n Aum Sri Ganesh !")
	fmt.Printf("\n %#04X", submod.CheckFunction())
	fmt.Printf("\n %#04d", submod.TestVariable)
	fmt.Print("\n Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v", name)
}
