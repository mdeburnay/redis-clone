package main

import (
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	// Start the server in a goroutine
	go main()

	// Give the server a chance to start up
	time.Sleep(100 * time.Millisecond)

	// Connect
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
}
