package main

import "fmt"

func main() {

	fmt.Println(" For Loop: ")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n While Loop: ")
	iter := 1
	for iter <= 10 {
		fmt.Println(iter)
		iter++
	}

	fmt.Println()
}
