package main

import "fmt"

func main() {
	// no default fall through need to explicitly say "fallthrough"
	// Multiple Evaluation Not Possible
	// OR by default using , between arguments

	fmt.Println(" Example for Testing Multiple Evalutation ")
	n := 23
	fmt.Println(" n =", 23)
	switch {
	case n < 40:
		fmt.Println(" n is less than 40")
	case n > 2:
		fmt.Println(" n is greater than 2")
	}

	fmt.Println(" Example for Testing Fallthrough ")
	switch {
	case n < 40:
		fmt.Println(" n is less than 40")
		fallthrough
	case n == 2:
		fmt.Println(" n is equal to 2 even though this is not true")
	}
}
