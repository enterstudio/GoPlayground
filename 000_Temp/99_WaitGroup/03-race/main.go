package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var inc int

func main() {
	inc = 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			m := inc
			runtime.Gosched()
			m++
			inc = m
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc)
}
