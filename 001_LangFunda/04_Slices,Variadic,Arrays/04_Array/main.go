package main

import "fmt"

func printit(str [3]string) {
	for k, v := range str {
		fmt.Println(k, v)
	}
}

func main() {
	//names := [4]string{"Ram","Sham","Vivek"} // Arrays are always fixed sized
	names := [3]string{"Ram", "Sham", "Vivek"}

	printit(names)
}
