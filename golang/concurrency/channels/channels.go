package channels

import (
	"fmt"
	"strings"
)

// shout has two parameters: a Receive-only channel ping and a send-only channel pong
// Note: The use of the "<-" in function signature. It simply takes whatever string it gets from the ping channel, converts it into uppercase and
// appends a few exclamation marks before forwarding the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		// Read from the ping channel.
		// Note: The goroutine waits here - it blocks until something is received on this channel
		// Follows the comma-ok syntax. ok is false when a channel is closed
		if s, ok := <-ping; ok {
			// Do something
			pong <- fmt.Sprintf("%s !!!!!", strings.ToUpper(s))
		}
	}
}

// Channels allow us to communicate between goroutines (pipes): Either get data from a goroutine or push data to a goroutine
func ChannelsDemo() {
	// Create two channels: ping is what we send to, pong is what we receive from
	ping := make(chan string)
	pong := make(chan string)

	// Start goroutine in the background
	// (If this is commented out, we will get an "all goroutines are asleep - deadlock!" error.
	// This is Go's way of telling that we are sending something to a channel but nothing is receiving it.)
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		// Print a prompt
		fmt.Print("-> ")
		// Get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToUpper(userInput) == "Q" {
			break
		}
		// Send userInput to the ping channel
		ping <- userInput

		// wait for a response
		response := <-pong
		fmt.Println("Response:", response)
	}

	fmt.Println("All done, closing channels!")
	close(ping)
	close(pong)
}
