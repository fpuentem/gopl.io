package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := fmt.Sprintf("http://")
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}

}

// WaitForServer attepts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attemps fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 1; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retriving...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
