package channels

import (
	"fmt"
	"sync"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

var (
	programRunning = true
	programMutex   sync.Mutex
)

func server1(ch chan<- string, quitChan <-chan struct{}) {
	for {
		programMutex.Lock()
		if !programRunning {
			programMutex.Unlock()
			return
		}
		programMutex.Unlock()

		// First time it sleeps for 6 seconds
		time.Sleep(6 * time.Second)

		select {
		case <-quitChan:
			fmt.Println("server1 stopped")
			return
		default:
			ch <- "This is from server 1"
		}
	}
}

func server2(ch chan<- string, quitChan <-chan struct{}) {
	for {
		programMutex.Lock()
		if !programRunning {
			programMutex.Unlock()
			return
		}
		programMutex.Unlock()

		// First time it sleeps for 3 seconds
		time.Sleep(3 * time.Second)

		select {
		case <-quitChan:
			fmt.Println("server2 stopped")
			return
		default:
			ch <- "This is from server 2"
		}
	}
}

func quitAfter(ch chan<- struct{}, seconds time.Duration) {
	time.Sleep(seconds * time.Second)
	// Quitting and closing channels
	fmt.Println("Quitting...")
	close(ch)
}

func ChannelSelectDemo() {
	std.PrintHeader("Select with Channels")

	// Create channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Channel to trigger a quit
	quitChan := make(chan struct{})

	go server1(channel1, quitChan)
	go server2(channel2, quitChan)

	// Quit after 15 seconds
	go quitAfter(quitChan, 15)

	// Receive from channels
	for {
		programMutex.Lock()
		if !programRunning {
			programMutex.Unlock()
			break
		}
		programMutex.Unlock()

		// Demonstrating that select statement, if there are multiple matching cases, unlike a switch-case, which would choose the first matching case,
		// The select statement being for channels, i.e. for concurrent programming, it chooses a case at random.
		select {
		case s1, ok := <-channel1:
			if ok {
				fmt.Println("Case one:", s1)
			}
		case s2, ok := <-channel1:
			if ok {
				fmt.Println("Case two:", s2)
			}
		case s3, ok := <-channel2:
			if ok {
				fmt.Println("Case three:", s3)
			}
		case s4, ok := <-channel2:
			if ok {
				fmt.Println("Case four:", s4)
			}
		// A quit channel helps quitting out of an endless loop / goroutine
		case <-quitChan:
			programMutex.Lock()
			programRunning = false
			programMutex.Unlock()
		}
		// default:
		// 	// useful for avoiding deadlock
		// }
	}

	// Close channels
	fmt.Println("Closing channels...")
	close(channel1)
	close(channel2)
}
