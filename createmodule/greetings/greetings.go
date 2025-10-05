/*
* https://go.dev/doc/tutorial/create-module
**/
package greetings

import "fmt"

/*
* A function whose name starts with a capital letter
* can be called by a function not in the same package.
* This is known in Go as an exported name. (https://go.dev/tour/basics/3)
**/
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
