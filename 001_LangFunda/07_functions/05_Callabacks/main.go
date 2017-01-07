package main

import "fmt"

func appointment(people map[string]int, checker func(int) bool) map[string]int {
	// Create the List of Appointments
	finallist := make(map[string]int, len(people))

	for key, value := range people {
		// Check if we can get appointments
		if checker(value) {
			finallist[key] = value + 1 // Update the Number of Appointments
		}
	}
	// Return the Final process list
	return finallist
}

func main() {

	fullList := map[string]int{
		"Eshwer": 2,
		"Mahesh": 5,
		"Rena":   4,
		"Preeti": 6,
	}

	fmt.Println(" Full List: ", fullList)

	// Function to Perform the Check on the Appointments
	apptChecker := func(n int) bool {
		// Cant have more than 4 appointments per person
		if n > 4 {
			return false
		}
		return true
	}

	listnew := appointment(fullList, apptChecker)
	fmt.Println(" Scheduled: ", listnew)
	listnew = appointment(listnew, apptChecker)
	fmt.Println(" Scheduled2: ", listnew)
}
