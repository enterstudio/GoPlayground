package main

import "fmt"

func main() {
	// Make is used to define the slices the correct way
	//   as directly creating slices have memory implications
	names := make([]string, 3, 5)
	// This defines a slice with length of 3 and capacity of 5
	fmt.Println("Capacity = ", cap(names), "Length =", len(names), "value= ", names)

	//Adding Values
	names[0] = "Adi"
	names[1] = "Rohan"
	names[2] = "Manish"
	fmt.Println("Capacity = ", cap(names), "Length =", len(names), "value= ", names)

	// Adding more values
	//names[3] = "Gaurav" // Wrong way to add values to a defined slice
	names = append(names, "Gaurav")
	fmt.Println("Capacity = ", cap(names), "Length =", len(names), "value= ", names)

	// Even more names
	names = append(names, "Jayesh")
	names = append(names, "Vipin")
	fmt.Println("Capacity = ", cap(names), "Length =", len(names), "value= ", names)
	// Notice that the Capacity gets doubled on every additional element

	names2 := []string {"Deven", "Jagan"}
	fmt.Println("Capacity = ", cap(names2), "Length =", len(names2), "value= ", names2)
	// Combining Slices
	// names = append(names, names2) // - This is wrong as we can add string slice directly
	names = append(names, names2...)// Representation in Variadic
	fmt.Println("Capacity = ", cap(names), "Length =", len(names), "value= ", names)
}
