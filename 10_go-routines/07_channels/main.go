package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var counter int
var c_chan = make(chan int)
var done_chan = make(chan bool)
var x int64

func main() {
	go f1("A1 :")
	go f1("A2 :")
	go puller()
	<-done_chan
	fmt.Println("Final Counter:", counter)
}

func f1(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		c_chan <- 1
		fmt.Println(s, i)
		if i == 9 {
			atomic.AddInt64(&x, 1)
			fmt.Println("XXXXXXXXX", x)
		}

		if atomic.LoadInt64(&x) == 2 {
			close(c_chan)
		}
	}
}

func puller() {
	for {
		i, more := <-c_chan
		if more {
			counter += i
			fmt.Println("Counter:", counter)
		} else {
			done_chan <- true
			close(done_chan)
			return
		}
	}
}
