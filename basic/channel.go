package main

import (
	"fmt"
	"time"
)

func UnbufferedChannel() {
	channel := make(chan string, 1)
	go func() {
		channel <- "hello world!"
	}()
	// We don't need to use a WaitGroup to sync the goroutine as
	// the defaul nature of channels is to block until data is received
	fmt.Println(<-channel)

	//for value := range channel {
	//	fmt.Printf("s: %s", value)
	//}

}

// Now, senders don't need to wait until some goroutine picks the data they are sending
// So the goroutine buffers the string into the channel and continues, no needing to waiting
// for the receiver
func BufferedChannel() {
	channel := make(chan string, 1)
	go func() {
		channel <- "Hello world! 1"
		//channel <- "Hello World! 2"
		println("Finishing unblocked goroutine")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println(<-channel)
}

// DIRECTIONAL CHANNELS
// When we use channels as parameters we restrict their directionality
// so that hey can be used only to send or to receive (static type check)

func DirectionalChannel() {
	channel := make(chan string, 1)

	// constraint of only input channel
	// go func(ch <- chan string){  // uncomment to see the type error
	go func(ch chan<- string) {
		ch <- "Hello"
		println("Finishing goroutine")
	}(channel)
	time.Sleep(time.Second)
	msg := <-channel
	fmt.Println(msg)
}

func main() {
	//UnbufferedChannel()
	//BufferedChannel()
	//DirectionalChannel()
}

//////////////////////////////////////////////
//  SUMMARY
//////////////////////////////////////////////
//  create a channel with make command
//  ex. make(chan int)

// send-only :  chan <- val int
// receive-only : val int <- chan

// sender side block untill receiver is available
// block receiver till msg is available
// can decouple sender and receiver with buffered channels