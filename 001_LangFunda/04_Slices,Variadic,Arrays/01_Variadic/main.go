package main

import "fmt"

// For Parameters the variadic ...Params
func greet(prefix string, s ...string) {
	fmt.Println(prefix)
	for key, value := range s {
		fmt.Print(" Hi There ", value, "(", key, ") \n")
	}
	fmt.Println()
}

func main() {
	fmt.Println(" Passic Variadic Arguments then <Arg>...")
	s := []string{"Avin", "Rohan"} // String Array
	greet("\n People Here: ", s...)
}
