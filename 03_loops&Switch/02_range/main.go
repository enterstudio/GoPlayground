package main

import "fmt"

func main() {
	val := "This is a String"

	fmt.Println("\n String Print for =", val)
	for key, value := range val {
		fmt.Printf("\n %2d - %4d - %s - %08b - %#02x", key, value, string(value), value, value)
	}

	fmt.Println()
}
