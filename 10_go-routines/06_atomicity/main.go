package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	wg.Add(2)
	go f1("A1")
	go f1("A2")
	wg.Wait()
}

func f1(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, "iter:", i, "Counter:", counter)
	}
	wg.Done()
}
