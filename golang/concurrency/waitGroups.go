package concurrency

import (
	"fmt"
	"sync"
)

var msg string
var wg = sync.WaitGroup{}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func WaitGroups() {

	// Challenge:
	// (1) Modify this code so that the calls to updateMessage() on lines 24, 27 and 30 run as goroutines
	// (2) Implement waitgroups so that the program runs properly, and print out three different messages in the exact same order.
	// (3) Write a test for all three functions in the program: `updateMessage()`, `printMessage()` and `ConcurrencyOperations()`

	msg = "Hello, world!"

	// Solution 1: Using wait group directly
	// wg.Add(1)
	// go updateMessage("Hello, universe!")
	// wg.Wait()
	// printMessage()

	// wg.Add(1)
	// go updateMessage("Hello, cosmos!")
	// wg.Wait()
	// printMessage()

	// wg.Add(1)
	// go updateMessage("Hello, world!")
	// wg.Wait()
	// printMessage()

	// Solution 2: Using a slice
	messages := []string{"Hello, universe!", "Hello, cosmos!", "Hello, world!"}

	for _, msg := range messages {
		wg.Add(1)
		go updateMessage(msg)
		wg.Wait()
		printMessage()
	}

}
