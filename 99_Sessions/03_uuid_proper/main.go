package main

import (
	"fmt"
)

func main() {

	fmt.Println("\nGenerate UUID (pure randome)                : ", GetUUID())
	fmt.Println("\nGenerate UUID (pure randome) in Bytes Array : ", GetUUIDbytes())
	uid, _ := DerivedUUID("Test1")
	fmt.Println("Generated UUID with Text(Test 1)   : ", uid)
	uid, _ = DerivedUUID("Test1")
	fmt.Println("Generated UUID with Text(Test 1)   : ", uid)
	uid, _ = DerivedUUID("Testing2")
	fmt.Println("Generated UUID with Text (Testing2): ", uid)
}
