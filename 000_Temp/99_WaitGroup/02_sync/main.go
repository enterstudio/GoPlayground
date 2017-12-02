/// Code to Perform Synchronization with wait group
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t:", runtime.GOOS)
	fmt.Println("Architecture\t:", runtime.GOARCH)
	fmt.Println("Number of CPU\t:", runtime.NumCPU())

	fmt.Println("Number of Goroutines\t:", runtime.NumGoroutine())
	wg.Add(2)
	go routine(22)
	go routine(42)
	wg.Wait()
	fmt.Println("Number of Goroutines\t:", runtime.NumGoroutine())
}

func routine(val int) {
	fmt.Println("Number of Goroutines\t:", runtime.NumGoroutine())
	fmt.Println("Value is", val)
	wg.Done()
}
