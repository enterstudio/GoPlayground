package main

import "fmt"

func main() {

	// Fully Dynamic MAP - reference type - MAPS are NOT THREAD SAFE
	students := make(map[uint32]string, 2) // Capacity is possible but grows
	students[2012] = "Adit"
	students[5245] = "Bharath"

	// Wrong way
	//fmt.Println(" cap=", cap(students), " len=", len(students), " value=", students)
	fmt.Println(" len=", len(students), " value=", students)

	//Direct Map definitions
	idm := map[string]string{
		"Read":   "Book1",
		"NoGood": "Book2",
	}
	fmt.Println(" len=", len(idm), " value=", idm)

	idm["VeryGood"] = "Book3"
	fmt.Println(" len=", len(idm), " value=", idm)

	//Check Existence
	if val, exists := idm["VeryGood"]; exists {
		delete(idm, "VeryGood")
		fmt.Println(" Existing Value=", val)
	} else {
		fmt.Println(" This does not exists ")
	}
	fmt.Println(" len=", len(idm), " value=", idm)
}
