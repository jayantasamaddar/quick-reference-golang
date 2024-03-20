package pc

import (
	"fmt"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

// Running total of successful pizzas made
var pizzasMade int

// Running total of Unsuccessful pizzas attempted
var pizzasFailed int

// Total pizzas attempted
var total int

// Struct for Producer
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		// Sets up a delay (for at least 1s): Simulate time taken to make a pizza
		delay := rand.Intn(5) + 1

		fmt.Printf("Received order #%d!\n", pizzaNumber)

		// Number that simulates that if this number hits, it fails
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza %d. It will take %d seconds...\n", pizzaNumber, delay)
		// Delay for a bit (artificial delay to simulate preparation)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for Pizza #%d! ***", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making Pizza #%d! ***", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza #%d is ready! ***", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *Producer) {
	// Keep track of which pizza we are trying to make
	var i int

	// Run forever or until we receive a quit notification from the quit channel

	// Try to make pizzas
	for {
		currentPizza := makePizza(i)
		//try to make a pizza

		if currentPizza != nil {
			// Set the new i
			i = currentPizza.pizzaNumber

			select {
			// we tried to make a pizza (we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}

		// Decision structure: Check whether the pizza is made,
	}
}

func PizzeriaProducerConsumerProblem() {
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Print out a message
	fmt.Println("The Pizzeria is open for business!")
	fmt.Println("----------------------------------")

	// Create a Producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// Run the Producer in the background
	go pizzeria(pizzaJob)

	// Create and run Consumer
	for order := range pizzaJob.data {
		if order.pizzaNumber <= NumberOfPizzas {
			if order.success {
				fmt.Println(order.message)
				fmt.Printf("Order #%d is out for delivery!\n", order.pizzaNumber)
			} else {
				fmt.Println(order.message)
				fmt.Println("The customer is really mad!")
			}
		} else {
			fmt.Println("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				fmt.Println("Error closing channel:", err)
			}
		}
	}

	// Print out the ending message
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("Done for the day!")

	fmt.Printf("We made %d pizzas, but failed to make %d, with %d attempts in total.\n", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		fmt.Println("It was an awful day...")
	case pizzasFailed >= 6:
		fmt.Println("It was not a very good day...")
	case pizzasFailed >= 4:
		fmt.Println("It was an okay day...")
	case pizzasFailed >= 2:
		fmt.Println("It was a pretty good day!")
	default:
		fmt.Println("it was a great day!")
	}
}
