package main

import (
	"fmt"
	"time"
)

func printMessage(message string, channel chan string) {
	for i := 0; i < 1; i++ {
		fmt.Println(message)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "done"
}

func main() {
	var channel chan string
	channel = make(chan string)
	go printMessage("hello from go routine", channel)

	response := <-channel

	fmt.Println(response)
	// go printMessage("go wow")
	// go printMessage("we miss classes")

}
