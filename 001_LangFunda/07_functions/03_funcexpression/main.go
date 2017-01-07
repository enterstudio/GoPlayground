package main

import "fmt"

func main() {
	// Creating Anonymous function with "func main" scope
	half := func(number int) (int, bool) {
		return number / 2, (number % 2) == 0
	}

	if hnum, iseven := half(22); iseven {
		fmt.Println(" Even =", hnum)
	} else {
		fmt.Println(" Odd=", hnum)
	}

}
