package main

import "fmt"

func idGenerator() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func greeting(name string) func() string {
	return func() string {
		return fmt.Sprint(" Hello ", name, " , Nice to meet you")
	}
}

func main() {
	idgen1 := idGenerator() // Independent closures
	idgen2 := idGenerator()

	fmt.Println(idgen1())
	fmt.Println(idgen1())
	fmt.Println(idgen1())

	fmt.Println(idgen2())
	fmt.Println(idgen1())
	fmt.Println(idgen2())

	greet1 := greeting("Jay") // Parameter based custom generators using closure
	fmt.Println(greet1())
	fmt.Printf("%T\n", greet1)
}
