package main

import (
	"testing"
)

// func TestServer(t *testing.T) {
// 	// Start the server in a goroutine
// 	go main()

// 	// Give the server a chance to start up
// 	time.Sleep(100 * time.Millisecond)

// 	// Connect
// 	conn, err := net.Dial("tcp", "127.0.0.1:6379")
// 	if err != nil {
// 		t.Logf("Could not connect to server: %v", err)
// 	}
// 	defer conn.Close()

// 	_, err = conn.Write([]byte("PING\r\n"))
// 	if err != nil {
// 		t.Fatalf("Could not write to server: %v", err)
// 	}

// 	buf := make([]byte, 1024)
// 	n, err := conn.Read(buf)
// 	if err != nil {
// 		t.Fatalf("Could not read from server: %v", err)
// 	}

// 	expected := "+OK\r\n"
// 	if string(buf[:n]) != expected {
// 		t.Fatalf("Expected %q, got %q", expected, string(buf[:n]))
// 	}
// }

/*
Purpose of this next function is to be able to understand the
format of the input we will receive in order to parse it.

For example, if we receive data from the client in the form of a RESP string,
like this: "*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"

We need to be able to read the first byte, which will tell us what type of
data we are receiving.
*/

func TestReader(t *testing.T) {
	input := "$5\r\nhello\r\n"

	result, error := BulkReader(input)
	if error != nil {
		t.Fatalf("Error: %v", error)
	}

	expected := "hello"
	if result != expected {
		t.Fatalf("Expected %q, got %q", expected, result)
	}

}
