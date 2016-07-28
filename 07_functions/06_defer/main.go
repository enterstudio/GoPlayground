package main

import "fmt"

func main() {
	defer fmt.Println(" Sri Ganeshay Namh!")
	h1()
	for i := 0; i < 0xFFFFFF; i++ {
		m := i * i
		fmt.Sprint(string(m))
	}
}

func h1() {
	fmt.Println("Aum")
}
