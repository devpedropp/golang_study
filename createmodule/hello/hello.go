/*
* https://go.dev/doc/tutorial/call-module-code
**/
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// an empty name will return error
	message, err := greetings.Hello("")
	// a name will return the greeting
	// message, err := greetings.Hello("Gladys")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
