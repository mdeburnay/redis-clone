package main

import "net"

func main() {
	// Start server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	// Accept connections
	conn, err := l.Accept()
	if err != nil {
		panic(err)
	}

	// Close the connection once finished
	defer conn.Close()
}
