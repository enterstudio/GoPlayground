package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{}

	fmt.Println(p1)
	p1.Age = 15
	fmt.Println(p1)
	p1.Name = "Takshak"
	fmt.Println(p1)

	p2 := new(Person)
	fmt.Println(p2)
	p2.Age = 25
	fmt.Println(p2)
	p2.Name = "Niki"
	fmt.Println(p2)
}
