package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Hari Aum!")
		wg.Done()
	}()

	go func() {
		runtime.Gosched()
		fmt.Println("Hare Kirshna!")
		wg.Done()
	}()
	wg.Wait()
}
