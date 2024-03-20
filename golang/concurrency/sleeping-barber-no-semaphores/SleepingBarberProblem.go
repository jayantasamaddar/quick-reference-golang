// This is a simple demonstration of how to solve the Sleeping Barber problem, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is open, and
// clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the waiting room for new clients,
// and if one or more is there, a haircut takes place. Otherwise, the barber goes to sleep until a new client arrives.
//
// The rules are as follows:
//
//   - If there are no customers, the barber falls asleep in the chair
//   - A customer must wake the barber when he's asleep
//   - If a customer arrives while the barber is working, the customer leaves if all chairs are occupied
//     and sits in an empty chair if it's available
//   - Shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is empty
//   - After the shop is closed and there are no clients left in the waiting area, the barbers go home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Djikstra
//
// The point of the problem and its solution, was to make clear that in a lot of cases, the use of semaphores (mutexes) is not needed.

package slb

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

// Variables
// Hard constants
const (
	SEATING_CAPACITY int = 5
	ARRIVAL_RATE         = 100
	CUT_DURATION         = 1000 * time.Millisecond
	TIME_OPEN            = 10 * time.Second
)

func SleepingBarberProblem() {
	// Seed our random number generator
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Print a welcome message
	std.PrintHeader("The Sleeping Barber Problem")

	// Create channels if we need any
	clientChan := make(chan string, SEATING_CAPACITY)
	doneChan := make(chan bool)

	// Create the data structures that represent the BarberShop
	shop := BarberShop{
		ShopCapacity:    SEATING_CAPACITY,
		HaircutDuration: CUT_DURATION,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}
	fmt.Println(std.PrintC(std.Green, "The shop is Open!"))

	// Add Barbers (Run each barber as a goroutine)
	shop.AddBarber("Frank")
	shop.AddBarber("Paul")
	shop.AddBarber("Gary")
	shop.AddBarber("Kelly")
	shop.AddBarber("Stephen")
	shop.AddBarber("Pam")

	// Start the BarberShop as a Goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		// The Barbershop must be open for at least the time defined in TIME_OPEN
		<-time.After(TIME_OPEN)
		shopClosing <- true
		shop.CloseShopForTheDay()
		closed <- true
	}()

	// Add clients
	i := 1
	go func() {
		for {
			// Get a random number with average arrival rate
			randomMilliseconds := random.Int() % (2 * ARRIVAL_RATE)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.AddClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// Block until the barbershop is closed
	<-closed
}
