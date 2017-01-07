package main

import (
	"fmt"
 "github.com/satori/go.uuid"
)

func main() {
	// Fixed UUID
	tempuuid:= "bb685881-2bb2-4367-b584-0ed58493d8c7"

	fmt.Println("\nGenerate UUID V4 (pure randome): ", uuid.NewV4())

	uid, err := uuid.FromString(tempuuid)
	if err != nil {
		fmt.Println("Error Could not process the UUID " , tempuuid)
	} else {
		fmt.Println("Read UUID from String          : ", tempuuid, " = ", uid)
	}

	name1 := "Test 1"
	fmt.Println("Constant Input Text: ", name1)
	fmt.Println("Generated UUID with NewV4 as input + Text(Test 1)         : ", uuid.NewV5(uuid.NewV4(), name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Test 1)  : ", uuid.NewV5(uid, name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Test 1)  : ", uuid.NewV5(uid, name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Testing2): ", uuid.NewV5(uid, "Testing2"))
}
