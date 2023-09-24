package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	conn.Write([]byte("220 FTP Server Ready\r\n"))

	// Create a bufio reader to read client commands
	reader := bufio.NewReader(conn)

	// Current working directory
	currentDir := "/"

	for {
		// Read a command from the client
		cmd, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
				return
			}
			fmt.Println("Error reading command:", err)
			return
		}

		// Remove leading/trailing whitespace and newline character
		cmd = strings.TrimSpace(cmd)

		// Split the command into parts (command and argument)
		parts := strings.SplitN(cmd, " ", 2)
		command := strings.ToUpper(parts[0])
		arg := ""
		if len(parts) > 1 {
			arg = parts[1]
		}

		// Process FTP commands
		switch command {
		case "CD":
			// Change directory command
			newDir := arg
			if newDir == "" {
				conn.Write([]byte("550 Missing directory argument\r\n"))
			} else {
				// Attempt to change the directory
				newPath := filepath.Join(currentDir, newDir)

				if _, err := os.Stat(newPath); os.IsNotExist(err) {
					// Directory does not exist
					conn.Write([]byte("550 Directory does not exist\r\n"))
				} else {
					// Directory exists, update the current working directory
					currentDir = newPath
					conn.Write([]byte("250 Directory changed\r\n"))
				}
			}
		case "LS":
			// List directory contents command
			files, err := os.ReadDir(currentDir)
			if err != nil {
				fmt.Println("Error listing directory:", err)
				conn.Write([]byte("550 Error listing directory\r\n"))
				return
			}

			// Create a response message with the list of files and directories
			var response strings.Builder
			response.WriteString("150 Here comes the directory listing.\r\n")

			for _, file := range files {
				response.WriteString(file.Name())
				if file.IsDir() {
					response.WriteString("/ (directory)\r\n")
				} else {
					response.WriteString(" (file)\r\n")
				}
			}

			response.WriteString("226 Directory send OK.\r\n")

			// Send the response message to the client
			conn.Write([]byte(response.String()))
		case "GET":
			// Get file contents command
			// TODO: Implement file retrieval logic here
			// Example: Open and read the requested file, then send its contents to the client
			conn.Write([]byte("250 Sending file\r\n"))
		case "CLOSE":
			// Close connection command
			conn.Write([]byte("221 Closing connection\r\n"))
			return
		default:
			// Invalid command
			conn.Write([]byte("500 Unknown command\r\n"))
		}
	}
}

// The rest of your code remains unchanged...
func main() {
	listen, err := net.Listen("tcp", ":21")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Println("FTP Server Listening on port 21")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go handleClient(conn)
	}
}
