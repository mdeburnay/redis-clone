package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

var ErrNoCommand = errors.New("Error: No command provided")

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		// ignore request and send back a PONG
		conn.Write([]byte("+OK\r\n"))
	}
}

func BulkReader(input string) (string, error) {
	reader := bufio.NewReader(strings.NewReader(input))

	// Read the first byte
	b, _ := reader.ReadByte()

	// If the first byte is a $, then we know we are receiving a bulk string
	if b != '$' {
		fmt.Println("Not a bulk string")
		return "", nil
	}

	// Next, read the number to determine the number of characters in a string
	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	// consume the \r\n
	reader.ReadByte()
	reader.ReadByte()

	output := make([]byte, strSize)
	reader.Read(output)

	return string(output), nil
}
