// Clock3 is a TCP server that periodically writes the time
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Parse command-line arguments
	port := flag.Int("port", 8000, "Port number")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle connection
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
