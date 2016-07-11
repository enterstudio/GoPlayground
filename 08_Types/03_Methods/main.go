package main

import "fmt"

type ExamInfo struct {
	RollNo   int
	Eligible bool
	Score int
}

type parentinfo struct{
	ParentName string
	homeAddress string
}

// Private/ Public is not restricted in the Promoted inner type, it only works for Package scope
type Student struct {
	ExamInfo
	parentinfo
	Name string
	Age  int
}

// Method Attached to "parentinfo" structure
//   - However Pass by Value does not change the underlying data
func (p parentinfo) UpdateInfo(parentName, address string) parentinfo {
	p.homeAddress = address
	p.ParentName = parentName
	return p
}

func (p* ExamInfo) UpdateEligibility(ExamScore int) bool {
	s := *p
	s.Score = ExamScore
	// Check the Eligibility
	s.Eligible = s.Score > 200
	*p = s
	return s.Eligible
}

func main() {

	s1 := Student{
		ExamInfo: ExamInfo{
			RollNo: 25468,
			Eligible: true,
			Score: 200,
		},
		parentinfo: parentinfo{
			ParentName: "Anjana",
			homeAddress: "22, Unknown Street",
		},
		Name: "Radhika",
		Age: 14,
	}

	fmt.Println(s1)
	s1.ParentName = "Deepak"
	fmt.Println(s1)
	s1.homeAddress = "45, New Street"
	fmt.Println(s1)
	// methods of the Promoted type can be accessed
	fmt.Println(s1.UpdateInfo("Kirti","101, Named Street"))
	fmt.Println(s1)
	// Updating Radhika's Latest Test scores
	fmt.Println(s1.UpdateEligibility(150))
	fmt.Println(s1)
}
