package main

import "fmt"

func main() {
	intStream := make(chan int)
	close(intStream)
	for i := 0; i < 5; i++ {
		integer, ok := <-intStream
		fmt.Printf("(%v): %v\n", ok, integer)
	}
}
