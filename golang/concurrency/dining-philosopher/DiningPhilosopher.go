package df

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

// Philosopher struct stores info about a philosopher
type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

// philosophers is a list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// Times philosopher has to eat before they are done
var hunger = 3

// Times philosopher takes to eat
var eatTime = 5000 * time.Millisecond

// Times philosopher takes to think
var thinkTime = 1000 * time.Millisecond

// Times philosopher takes to relax
var sleepTime = 500 * time.Millisecond

var orderMutex = &sync.Mutex{}
var orderOfLeaving = []string{}

func DiningPhilosopher() {
	// Print a welcome message
	fmt.Println(std.PrintC(std.Purple, "-------------------------------------------------------------------------------------------"))
	fmt.Println(std.PrintC(std.Purple, "Dining Philosophers' Problem"))
	fmt.Println(std.PrintC(std.Purple, "-------------------------------------------------------------------------------------------"))

	fmt.Println(std.PrintC(std.Yellow, "The table is empty!"))

	// Start the meal
	dine()

	// Print out a finish message
	fmt.Println(std.PrintC(std.Yellow, "The table is empty!"))

}

func dine() {
	// Controls dining
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// Controls seating
	seatedWg := &sync.WaitGroup{}
	seatedWg.Add(len(philosophers))

	// forks
	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// Start the meal
	for i := 0; i < len(philosophers); i++ {
		// Fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seatedWg)
	}
	wg.Wait()

	// Implement order the philosophers left
	fmt.Println("Order the philosophers left:", strings.Join(orderOfLeaving, ", "))
}

// diningProblem is the function fired off for each of our philosophers.
// It takes one philopher, waitGroup to determine when everyone is done, a map containing mutexes mapped to every fork,
// and a wait group to pause execution of every instance of this goroutine until everyone is seated at the table
func diningProblem(p Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seatedWg *sync.WaitGroup) {
	defer wg.Done()

	// Seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", p.name)
	seatedWg.Done()

	// Wait until everyone is seated.
	seatedWg.Wait()

	// Eat three times
	for i := hunger; i > 0; i-- {
		// Get a lock on the left and right forks. We have to choose the lower numbered fork first in order
		// to avoid a logical race condition, which is not detected by the -race flag in tests; if we don't do this,
		// we have the potential for a deadlock, since two philosophers will wait endlessly for the same fork.
		// Note that the goroutine will block (pause) until it gets a lock on both the right and left forks.
		if p.leftFork > p.rightFork {
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", p.name)
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", p.name)
		} else {
			forks[p.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", p.name)
			forks[p.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", p.name)
		}

		fmt.Printf("\t%s has both forks, and is eating...\n", p.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking...\n", p.name)
		time.Sleep(thinkTime)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()
		fmt.Printf("\t%s has put down the forks...\n", p.name)
	}

	fmt.Println(p.name, " is satisfied!")
	fmt.Println(p.name, " has left the table!")

	// Implement order the philosophers left
	orderMutex.Lock()
	orderOfLeaving = append(orderOfLeaving, p.name)
	orderMutex.Unlock()
}
