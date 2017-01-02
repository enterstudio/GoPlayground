package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
This File demonstrates the use of Trapping the SIGINT
or the Ctrl+C thing
*/

// CleanupFn is a Data type of Clean-up function
type CleanupFn func()

func cleanupFn() {
	fmt.Println("cleanup")
}

// SetIntrCleanup Function used to configure the catching of SIGTERM signal
//  from OS and add the respective cleanup-function and
//  proper exit status rather than the default one
func SetIntrCleanup(cfn CleanupFn, exitStatus int) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cfn()
		os.Exit(exitStatus)
	}()
}

func main() {

	SetIntrCleanup(cleanupFn, 0)

	for {
		fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}

}
