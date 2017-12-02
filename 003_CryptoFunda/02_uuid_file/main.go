package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
)

func main() {
	tempuuid := "bb685881-2bb2-4367-b584-0ed58493d8c7"
	// Fixed UUID
	fixuuid, err := ioutil.ReadFile(".uuid-fix")
	if err != nil {
		fmt.Println("Creating UUID File ...")
		fixuuid = []byte(uuid.NewV4().String())
		err = ioutil.WriteFile(".uuid-fix", fixuuid, 0644)
		if err != nil {
			fmt.Println("Error Could not write the UUID file ")
		}
	} else {
		fmt.Println("Found the Fixed UUID File")
	}

	fmt.Println("Fixed UUID: ", string(fixuuid))

	fmt.Println("\nGenerate UUID V4 (pure randome): ", uuid.NewV4())
	uid, _ := uuid.FromString(tempuuid)

	name1 := "Test 1"
	fmt.Println("Constant Input Text: ", name1)
	fmt.Println("Generated UUID with NewV4 as input + Text(Test 1)         : ", uuid.NewV5(uuid.NewV4(), name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Test 1)  : ", uuid.NewV5(uid, name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Test 1)  : ", uuid.NewV5(uid, name1))
	fmt.Println("Generated UUID with String UUID as input + Text (Testing2): ", uuid.NewV5(uid, "Testing2"))
}
