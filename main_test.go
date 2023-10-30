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
		t.Logf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("PING\r\n"))
	if err != nil {
		t.Fatalf("Could not write to server: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Could not read from server: %v", err)
	}

	expected := "+OK\r\n"
	if string(buf[:n]) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(buf[:n]))
	}
}
