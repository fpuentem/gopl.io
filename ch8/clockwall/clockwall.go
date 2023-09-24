// clockwall is a program that acts as a client of several clock servers.
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()
	locations := flag.Args() // Get the list of server locations
	// fmt.Println(locations)
	// fmt.Printf("%T\n", locations)
	// fmt.Printf("%T\n", locations[0])
	for _, loc := range locations {
		// fmt.Printf("%T\n", loc)
		// fmt.Println(loc)
		// Parse the server location (format: "City=host:port")
		var city, address string
		// loc = strings.TrimSpace(loc)
		_, err := fmt.Sscanf(loc, "%s %s", &city, &address)
		// fmt.Println(err)
		if err != nil {
			// fmt.Fprintf(os.Stderr, "Invalid location format: %s\n", loc)
			// os.Exit(1)
			panic(err)
		}

		// Connect to the clock2 server
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error connecting to %s: %v\n", city, err)
			continue
		}
		defer conn.Close()

		// Read and display the time from the server
		buffer := make([]byte, 64)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading time from %s: %v\n", city, err)
			continue
		}

		fmt.Printf("%s: %s", city, buffer)
	}
}
