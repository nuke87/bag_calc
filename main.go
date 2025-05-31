package main

import "fmt"

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan string)
	ch1 := "hi"
	x := 1

	// Start a goroutine that sends a message to the channel
	go func() {
		ch <- "Hello from the goroutine!"
	}()

	// Receive the message from the channel
	message := <-ch

	// Print the received message
	fmt.Println(message)
	fmt.Println(ch1)
	fmt.Println(x)
}

//das ist ein Kommentar, um git zu testen
// noch ein Kommentar als test
// und noch ein Kommentar
// und noch ein Kommentar
