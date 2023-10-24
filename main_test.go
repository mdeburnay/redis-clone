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
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	for {
// 		buf := make([]byte, 1024)

// 		_, err = conn.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		conn.Write([]byte("+OK\r\n"))
// 	}
// }

/*
	FOR TESTING REFERENCE:
	For Simple Strings, the first byte of the reply is "+"
	For Errors, the first byte of the reply is "-"
	For Integers, the first byte of the reply is ":"
	For Bulk Strings, the first byte of the reply is "$"
	For Arrays, the first byte of the reply is "*"
*/

/*
	1. Take in a slice of strings
	2. As it is a string parse the data in accordance with simple string RESP protocol
	3. Return a string and an error
*/

func TestHandleSetCommand(t *testing.T) {
	args := []string{"SET", "name", "john"}

	expected := "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$4\r\njohn"

	actual, err := HandleSetCommand(args)

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}
