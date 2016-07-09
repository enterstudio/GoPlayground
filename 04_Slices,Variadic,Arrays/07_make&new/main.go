package main

import "fmt"

func main() {

	// Creating arrays using different techniques

	//var age []int = new([]int) // - Wrong as New Returns a Pointer Zero valued object of the given Type
	var age *[]int = new([]int)
	fmt.Println(age)
	fmt.Println(*age)
	if *age == nil {
		fmt.Println(" So New Creates a Zero Length ")
	}

	var tem []int = make([]int, 5, 20) // Size 5 all will be Zero Value and total capacity 20
	fmt.Println(tem)
	//fmt.Println(*tem) // Wrong as the Array is not a pointer by default like in C

	// INITIALIZED data and always with Data structures = make
	// ZEROED data and pointer on All TYPES = new
	var md []int = make([]int, 0) // Similar to New but its an actual value not Pointer
	fmt.Println(md)
	if md == nil {
		fmt.Println(" Make with Zero size is NIL ")
	}
}
