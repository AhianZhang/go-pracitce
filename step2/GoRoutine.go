package main

import "fmt"

func main() {
	go func() {
		defer fmt.Println("A.defer")
		go func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {

	}
}
