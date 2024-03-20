package concurrency

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func MutexDemo() {
	// Variable for Bank Balance
	var bankBalance int

	// Define mutex
	var balance sync.Mutex

	// Print starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	// Define weekly revenue
	incomes := []Income{
		{"Main Job", 500},
		{"Gifts", 10},
		{"Part-time Job", 50},
		{"Investments", 100},
	}

	wg.Add(len(incomes))
	// Loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}

	wg.Wait()

	// Print the final balance
	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
}
