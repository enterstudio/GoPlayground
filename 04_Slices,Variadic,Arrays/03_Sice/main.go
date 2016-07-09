package main

import "fmt"

func main() {
	Myslice := []string{"S", "T", "P", "L", "K"}
	//slice := []string{"S", "T", "P", 'L', 'K'} // - This is wrong as rune cant be part of String slice
	fmt.Println(Myslice)
	fmt.Println(Myslice[2])
	fmt.Println(Myslice[2:4])
	fmt.Println("MySlice"[3:5])
}
