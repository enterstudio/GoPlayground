package main

import "fmt"

func main() {
	s := "This is a String"
	fmt.Println("String = ", s)
	// Try out Immutability
	//s[0] = 'k' // This does not work
	//fmt.Println("String = ", s)

	// Create a Slice - They are mutable
	fmt.Println("String[:4] = ", s[:4])
	fmt.Println("String[1:4] = ", s[1:4])

	// Conversion T(x) where T is the type and x is the input variable
	fmt.Println(" Converted to Byte slice ", []byte(s))
	fmt.Println(" ByteSice to String ", string([]byte{'h', 'e', 'l', 'l', 'o'}))

	// Single Character in Go is RUNE or int32 (UTF-8)
	fmt.Printf(" Printing the 'a' Type: %T \n", 'a')
	fmt.Printf(" Printing the Type for rune(0x55) %T \n", rune(0x55)) // 'U'
}
