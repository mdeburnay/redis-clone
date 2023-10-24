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

func TestHandleSetCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "bulk strings",
			args:     []string{"SET", "name", "john"},
			expected: "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$4\r\njohn",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := HandleSetCommand(test.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if actual != test.expected {
				t.Fatalf("expected %q, but got %q", test.expected, actual)
			}
		})
	}
}
