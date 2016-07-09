package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10

	fmt.Println("\n Display the Value:")
	fmt.Println("a =", a)
	fmt.Println("\n Display the Address:")
	fmt.Println("&a =", &a)

	b := &a
	fmt.Println("\n Print the Assigned address")
	fmt.Println(b)
	fmt.Println("\n Print the actual type of the variable")
	fmt.Printf("%T\n", b)
	fmt.Println("\n Print the type of the variable using Reflect package")
	fmt.Println(reflect.TypeOf(b))
	fmt.Println("\n Value in b =", *b)

	fmt.Println("\n Change the Value using Pointers '*b = 32' ")
	*b = 32 // Change the Value stored at the address pointed by 'b'
	fmt.Println(" a =", a)

	fmt.Println()
}
