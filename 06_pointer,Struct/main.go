package main

import "fmt"

// Defining a Struct
//  Small case in name means its Private
//  Same rules goes for members of the Struct
type person struct {
	name string
	age int
}
func main() {

	p1:= person{"Amit", 21}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p1.name) // Reflection in Action

	p2:= new(person) // Creating using New - Should be a Pointer
	p2.name = "Vivek"
	fmt.Println(p2) // Note this is Pointer but is implicitly access
	fmt.Println(*p2) // With converted access it still looks same as normal definition

	//*p2.age = 55// Wrong this would not work as '*' operator works only for value types
	// Struct is a Reference type
	p2.age = 55
	fmt.Println(p2) // Should be able to do Pointer based Assignment
}
