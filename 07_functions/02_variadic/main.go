package main

import "fmt"

func maxout(inlist ...int) int {
	var large int
	for _, n := range inlist {
		if n > large {
			large = n
		}
	}
	return large
}
func main() {
	fmt.Println("Largest number =", maxout(31, 2, 45, 1, 22, 32, 42, 50, 1, 2, 6))
}
