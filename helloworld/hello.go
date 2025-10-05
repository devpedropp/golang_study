/*
https://go.dev/doc/tutorial/getting-started

1. run "go mod init example/hello"
2. go run .
3. search packages on pkg.go.dev
4. run "go mod tidy" to download imported module
*/

package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
