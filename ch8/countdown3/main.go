package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//  ... create abort channel ...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // Read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
}

func launch() {
	fmt.Println("...countdown zero...")
}
