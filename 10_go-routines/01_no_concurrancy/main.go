package main

import "fmt"

func main() {
	f1("A1")
	f1("A2")
}

func f1(s string) {
	for i := 0; i < 40; i++ {
		fmt.Println("f1:", s, "Value:", i)
	}
}
