package main

import (
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Create a pipe to capture the standard output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	oldStdout := os.Stdout

	// Redirect os.Stdout to the pipe
	os.Stdout = w
	defer func() {
		// Restore the original os.Stdout and close the pipe
		os.Stdout = oldStdout
		w.Close()
	}()

	// Call the function that prints "Hello World"
	HelloWorld()

	// Read from the pipe to capture the output
	output := make([]byte, 1024)
	n, err := r.Read(output)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the captured output matches the expected string
	expected := "Hello world!\n"
	actual := string(output[:n])
	if actual != expected {
		t.Errorf("Expected: %s\nActual: %s", expected, actual)
	}
}

// Need to check how we will handle multiple strings ie. Hello World rather than just Hello
func TestSerialiseInput(t *testing.T) {
	result, err := SerializeInput("SET", "name", "john")

	// Check for errors
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Check the result against the expected value
	expected := "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$4\r\njohn"
	if result != expected {
		t.Errorf("Expected: %s\nActual: %s", expected, result)
	}
}
