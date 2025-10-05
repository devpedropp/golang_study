/*
* https://go.dev/doc/tutorial/create-module
* https://go.dev/doc/tutorial/handle-errors
**/
package greetings

import (
	"errors"
	"fmt"
)

/*
* A function whose name starts with a capital letter
* can be called by a function not in the same package.
* This is known in Go as an exported name. (https://go.dev/tour/basics/3)
*
* Any Go function can return multiple values. (https://go.dev/doc/effective_go#multiple-returns)
**/
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	// nil, means "no error"
	return message, nil
}
