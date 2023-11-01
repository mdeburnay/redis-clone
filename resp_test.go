package main

import (
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	input := "SET mykey myvalue\r\n"
	resp := NewResp(strings.NewReader(input))

	line, n, err := resp.readLine()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedLine := "SET mykey myvalue"
	if string(line) != expectedLine {
		t.Errorf("Expected line %q, got %q", expectedLine, line)
	}

	expectedN := len("SET mykey myvalue\r\n")
	if n != expectedN {
		t.Errorf("Expected n to be %d, got %d", expectedN, n)
	}
}
