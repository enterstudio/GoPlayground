package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go f1("A1")
	go f1("A2")
	wg.Wait()
}

func f1(s string) {
	for i := 0; i < 40; i++ {
		fmt.Println("f1:", s, "Value:", i)
	}
	wg.Done()
}
