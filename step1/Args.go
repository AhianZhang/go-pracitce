package main

import (
	"fmt"
	"os"
)

/**
  go run ./Args.go Hello World
*/
func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("your input keyword:", s)

}
