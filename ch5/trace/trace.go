package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("BigSlowOperation")()
	// ... lots of work ...
	time.Sleep(10 * time.Second) // simulate an slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
