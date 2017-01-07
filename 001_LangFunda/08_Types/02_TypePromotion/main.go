package main

import "fmt"

type ExamInfo struct {
	RollNo   int
	Eligible bool
}

// Inner Type Promotion is a example of the Go way of Inheritance that's purely data related
//   With names staring with small letter can still be private and with Caps are public
//   The members of the Inner promoted type can be directly accessed in the new struct
type Student struct {
	ExamInfo
	Name string
	Age  int
}

func main() {

	s1 := Student{
		ExamInfo: ExamInfo{
			22323,
			true,
		},
		Name: "Rohit",
		Age:  12,
	}

	fmt.Println(s1)
	fmt.Println(s1.RollNo)
	fmt.Println(s1.ExamInfo.RollNo)
	fmt.Printf("%T\n", s1.ExamInfo)
	fmt.Printf("%T\n", s1.RollNo)
}
