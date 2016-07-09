package main

import (
	"./submod"
	"fmt"
)

// checkker := true // This is wrong as the ':=' can only be used in functions

func Onefunc() {
	testvar := 35
	fmt.Println("Printing an Internal Var: ", testvar)
}

func main() {
	var name string
	fmt.Printf("\n Aum Sri Ganesh !")
	fmt.Println("\n Checking an Imported Submodule Public Function [submod.CheckFunction()] ")
	fmt.Printf(" %#04X\n", submod.CheckFunction())
	fmt.Println("\n Checking an Imported Submodule Public variable [submod.TestVariable] ")
	fmt.Printf(" %#04d\n", submod.TestVariable)
	fmt.Println("\n Checking Input From Terminal ")
	fmt.Print(" Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Printf("\n Hello %v\n", name)
	fmt.Println("\n Checking Function call and internal Variable definitions ")
	Onefunc()
}
