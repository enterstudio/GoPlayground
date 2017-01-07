package main

import "fmt"

func inconv(num int) (int, bool) {
	if num%2 == 0 {
		return num / 2, true
	}
	return num / 2, false
}
func main() {
	if half, iseven := inconv(52); iseven {
		fmt.Println(" Even =", half)
	} else {
		fmt.Println(" Odd =", half)
	}

	if half, iseven := inconv(49); iseven {
		fmt.Println(" Even =", half)
	} else {
		fmt.Println(" Odd =", half)
	}
}
