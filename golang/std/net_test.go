package std

import (
	"fmt"
	"io"
	"net"
	"testing"
)

const (
	PORT int = 3000
)

var address = fmt.Sprintf("%s:%d", "127.0.0.1", PORT)

// Create a TCP Server
func TestListener(t *testing.T) {
	done := make(chan struct{})
	// (1) Create a TCP Server capable of listening for incoming connections (called a listener).
	// The listener is not yet listening, only created.
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Bound to %q", listener.Addr())

	// (2) Spin off the listener in a goroutine to mimic server behavior by accepting connections concurrently with the client operations in this test.
	go func() {
		defer func() { done <- struct{}{} }()
		for {
			// Start the listener and wait for the next connection. Running this in an endless for loop
			// basically keeps the listener accepting new connections
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			// Handle the received connection in a new goroutine.
			go func(c net.Conn) {
				// Ensure the connection is closed and a message is sent to the 'done' channel when this goroutine completes.
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				// Simulate reading data from the connection.
				buf := make([]byte, 1024)
				for {
					// Reads data from the connection
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err) // log error if its not EOF
						}
						// After receiving the FIN packet, the Read method returns the io.EOF error,
						// indicating to the listener's code that the client closed that side of the connection.
						// Exit the goroutine on read error (including EOF)
						t.Log(err.Error()) // Likely EOF
						return
					}

					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	// Client side operation: Connect to the server
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Connected to TCP Server: %q!\n", address)

	// Close the client connection to the TCP Server to simulate end of communication
	conn.Close()
	t.Logf("Connection to TCP Server: %q closed!\n", listener.Addr().String())

	// Block until the first done channel message is received (i.e. from the connection handler)
	<-done

	// Close the TCP Server (Listener) after all connections are handled.
	t.Logf("Closing Listener at %q!", address)
	if err := listener.Close(); err != nil {
		t.Error("Error closing listener:", err)
	}
	t.Logf("Listener at %q closed!", address)

	// Block until the goroutine that runs the listener sends a done channel message
	<-done

	// Close the channel
	close(done)
}
