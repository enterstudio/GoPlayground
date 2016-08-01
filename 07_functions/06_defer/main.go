package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println(" Sri Ganeshay Namh!")
	h1()
	time.Sleep(time.Duration(2) * time.Second)
}

func h1() {
	fmt.Println("Aum")
}
