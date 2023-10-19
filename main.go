package main

import (
	"fmt"
	"reflect"
)

func HelloWorld() {
	fmt.Println("Hello world!")
}

/*
TODO: Build a function that will take in command arguments and send them to the correct function.
In Redis, it parses information like so:
- For Simple Strings, the first byte of the reply is "+"
- For Errors, the first byte of the reply is "-"
- For Integers, the first byte of the reply is ":"
- For Bulk Strings, the first byte of the reply is "$"
- For Arrays, the first byte of the reply is "*"
For example if we pass in a command such as "SET name max", this would be an array of three items, all of which are strings.
*/
func SerializeInput(command string, key string, value string) (string, error) {
	var formattedValue string
	stringPrefix := "*3\r\n"
	typeOfValue := reflect.TypeOf(value).Kind()
	if typeOfValue == reflect.String {
		formattedValue = stringPrefix + value
	}
	return "Hello", nil
}

func main() {
	resp, err := SerializeInput("SET", "name", "john")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
