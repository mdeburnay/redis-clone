package main

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

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

var ErrNoCommand = errors.New("Error: No command provided")

func HandleInput(args []string) (string, error) {
	if len(args) < 1 {
		return "", ErrNoCommand
	}

	var builder strings.Builder
	builder.WriteString("*")
	builder.WriteString(strconv.Itoa(len(args)))
	builder.WriteString("\r\n")

	for i, arg := range args {
		builder.WriteString("$")
		builder.WriteString(strconv.Itoa(len(arg)))
		builder.WriteString("\r\n")
		builder.WriteString(arg)
		if i < len(args)-1 {
			builder.WriteString("\r\n")
		}
	}

	return builder.String(), nil
}
