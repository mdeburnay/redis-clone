package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
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

func HandleInput(input string) (string, error) {

	// If there is no input, return an error
	if len(input) < 1 {
		return "", ErrNoCommand
	}

	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// consume /r/n
	reader.ReadByte()
	reader.ReadByte()

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))

	return "bop"
}
