package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex // <1>
	var data int
	go func() {
		memoryAccess.Lock() // <2>
		data++
		memoryAccess.Unlock() // <3>
	}()
	memoryAccess.Lock() // <4>
	if data == 0 {
		fmt.Println("the value is 0")
	} else {
		fmt.Printf("the value is %v\n", data)
	}
	memoryAccess.Unlock() // <5>
}
