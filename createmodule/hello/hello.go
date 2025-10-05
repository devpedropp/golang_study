/*
* https://go.dev/doc/tutorial/call-module-code
**/
package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
