package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hello!")
	}() // <1>
}
