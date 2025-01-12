package main

func main() {
	var receiveChan <-chan interface{}
	var sendChan chan<- interface{}
	dataStream := make(chan interface{})

	receiveChan = dataStream
	sendChan = dataStream
}
