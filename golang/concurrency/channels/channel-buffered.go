package channels

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// Print a got data message
		if i, ok := <-ch; ok {
			fmt.Println("Got", i, "from channel")
		}

		// Simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func ChannelBufferedDemo() {
	// Buffered channel with a fixed size.
	// Useful when:
	// -----------
	// (1) when you know how many goroutines you have launched
	// (2) Limit the number of goroutines launched
	// (3) Limit the amount of work that is queued up
	ch := make(chan int, 80)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("Sending %d to channel...\n", i)
		ch <- i
		fmt.Printf("Sent %d to channel...\n", i)
	}

	fmt.Println("Done! Closing channel...")
	close(ch)
}
