package sb

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

// Hard constants
const (
	AVAILABLE_SEATS int = 5
	DAILY_CUSTOMERS int = 25
	CUTTING_TIME        = 5 * time.Second
)

var SHOP_SERVICES = []string{"Haircut", "Shave", "Facial", "Bleaching", "Massage", "Hair Colouring"}

// Shop Events
const (
	EVENT_SHOP_OPEN int = iota
	EVENT_CUSTOMER_ARRIVAL
	EVENT_CUSTOMER_DEPARTURE
	EVENT_SEATING_VACANT
	EVENT_JOB_COMPLETION
	EVENT_SHOP_CLOSE
)

type BarberShop struct {
	barbers []*Barber
	seats   map[int]*Customer
	isOpen  bool
	jobs    Jobs
	events  chan Event
}

type Barber struct {
	name       string
	isSleeping bool
}

type Customer struct {
	customerNumber int
	requirement    string
}

type Job struct {
	jobNumber   int
	requirement string
	customer    *Customer
	barber      *Barber
}

type Jobs struct {
	current   []*Job
	completed []*Job
}

// Complete a job
func (jobs *Jobs) completeJob(j *Job, mutex *sync.Mutex) {
	mutex.Lock()
	for i, job := range jobs.current {
		if job.jobNumber == j.jobNumber {
			jobs.current = append(jobs.current[:i], jobs.current[i+1:]...)
			jobs.completed = append(jobs.completed, j)
			fmt.Println(std.PrintC(std.Green, fmt.Sprintf("Job #%d completed!", j.jobNumber)))
			// Check status and sleep barber
			j.barber.isSleeping = true
			mutex.Unlock()
			return
		}
	}
}

type Event struct {
	name    int
	message string
	data    interface{}
}

var jobMutex = &sync.Mutex{}
var seatsMutex = &sync.Mutex{}
var barbersMutex = &sync.Mutex{}
var shopMutex = &sync.Mutex{}

var shopWg = &sync.WaitGroup{}
var haircutWg = &sync.WaitGroup{}
var closeShopChan = make(chan interface{})

// Open shop
func (shop *BarberShop) Open() {
	// Open shop
	shop.isOpen = true
	fmt.Println(std.PrintC(std.Purple, "-----------------------------------------------------------------------------------"))
	fmt.Println(std.PrintC(std.Purple, "SHOP OPEN"))
	fmt.Println(std.PrintC(std.Purple, "-----------------------------------------------------------------------------------"))
	shop.events <- Event{name: EVENT_SHOP_OPEN}
}

// Handle BarberShop events: Events listen for BarberShop Activities
func (shop *BarberShop) HandleBarberShopEvents() {
LOOP:
	for {
		select {
		case event := <-shop.events:
			switch event.name {
			case EVENT_SHOP_OPEN:
				// Perform BarberShop Activities
				go shop.PerformShopActivities()
			case EVENT_SEATING_VACANT:
				barbersMutex.Lock()
				if barber := shop.getNextAwakeBarber(); barber != nil {
					shop.putBarberToSleep(barber)
				}
				barbersMutex.Unlock()

			case EVENT_SHOP_CLOSE:
				shop.isOpen = false
				break LOOP
			}
		case <-closeShopChan:
			shop.isOpen = false
			break LOOP
		}
	}
}

// Put barber to sleep
func (shop *BarberShop) putBarberToSleep(barber *Barber) {
	barber.isSleeping = true
	fmt.Println(std.PrintC(std.Blue, fmt.Sprintf("Barber: %s is sleeping", barber.name)))
}

// Find the next awake barber. Returns nil if not found.
func (shop *BarberShop) getNextAwakeBarber() *Barber {
	for _, barber := range shop.barbers {
		if !barber.isSleeping {
			return barber
		}
	}
	return nil
}

// Find the next available barber. Returns nil if not found.
func (shop *BarberShop) getNextAvailableBarber() *Barber {
	for _, barber := range shop.barbers {
		if barber.isSleeping {
			barber.isSleeping = false
			return barber
		}
	}
	return nil
}

// Pick the first available Customer. Returns nil if not found.
func (shop *BarberShop) pickFirstCustomer() *Customer {
	if len(shop.seats) > 0 {
		for key, cust := range shop.seats {
			tempCust := cust
			delete(shop.seats, key)
			return tempCust
		}
	}
	return nil
}

// Start a Job
func (shop *BarberShop) ExecuteJob(j *Job, mutex *sync.Mutex) {
	defer shopWg.Done()
	// Starting Job
	fmt.Printf("Starting Job #%d...\n", j.jobNumber)
	time.Sleep(1 * time.Second)

	// Running Job
	fmt.Println(std.PrintC(std.Blue, fmt.Sprintf("%s Job (#%d) started by Barber: %s for Customer: #%d...", j.requirement, j.jobNumber, j.barber.name, j.customer.customerNumber)))
	time.Sleep(CUTTING_TIME)

	// Complete Job
	shop.jobs.completeJob(j, mutex)
	// After Customer is done cutting hair make them exit
	shop.handleCustomerExit(j.customer, seatsMutex)
}

// Perform the BarberShop Activities
func (shop *BarberShop) PerformShopActivities() {
	// Until shop is open perform shop activities
	for shop.isOpen {
		// Look for available customers
		if len(shop.seats) > 0 {
			// Check if barber is available
			if barber := shop.getNextAvailableBarber(); barber != nil {
				// Get a customer from the seating
				if customer := shop.pickFirstCustomer(); customer != nil {
					// Start a Job: Assign Barber to customer and start fulfilling the requirement
					job := &Job{
						barber:      barber,
						jobNumber:   rand.Intn(1000000),
						requirement: customer.requirement,
						customer:    customer,
					}
					shopMutex.Lock()
					shop.jobs.current = append(shop.jobs.current, job)
					shopMutex.Unlock()
					shopWg.Add(1)
					go shop.ExecuteJob(job, jobMutex)
				}
			} else {
				fmt.Println(std.PrintC(std.Yellow, fmt.Sprintf("%d seated customers are waiting for an available barber...", len(shop.seats))))
				time.Sleep(1 * time.Second)

			}
		} else {
			fmt.Println(std.PrintC(std.Yellow, "Waiting for customers..."))
			time.Sleep(1 * time.Second)
		}
	}
	shopWg.Wait()
}

// Handle the entry of a customer
func (shop *BarberShop) handleCustomerEntry(cust *Customer, mutex *sync.Mutex) {
	// Try to accomodate customer
	fmt.Println("Trying to accomodate customer...")
	time.Sleep(1 * time.Second)

	mutex.Lock()
	// If the current seating is equal to Available seats
	if len(shop.seats) >= AVAILABLE_SEATS {
		// Ensure customer leaves
		fmt.Println(std.PrintC(std.Red, fmt.Sprintf("Customer #%d has left because waiting room is full!", cust.customerNumber)))
		return
	} else {
		// Otherwise add customer to seat
		shop.seats[cust.customerNumber] = cust
		fmt.Println(std.PrintC(std.Cyan, fmt.Sprintf("Added customer #%d to seating. Current seats: %d.", cust.customerNumber, len(shop.seats))))
	}
	mutex.Unlock()
}

// Handle the exit of a customer
func (shop *BarberShop) handleCustomerExit(cust *Customer, mutex *sync.Mutex) {
	mutex.Lock()
	delete(shop.seats, cust.customerNumber)
	if len(shop.seats) == len(shop.barbers)-1 {
		shop.events <- Event{name: EVENT_SEATING_VACANT}
	}
	mutex.Unlock()
}

// The main Sleeping Barber Function
func SleepingBarber() {
	// Initialize Barbershop
	barber1 := &Barber{
		name:       "John Legend",
		isSleeping: true,
	}
	shop := &BarberShop{
		barbers: []*Barber{barber1},
		seats:   make(map[int]*Customer),
		isOpen:  false,
		jobs: Jobs{
			current:   []*Job{},
			completed: []*Job{},
		},
		events: make(chan Event, 3),
	}

	// Handle BarberShop Events
	go shop.HandleBarberShopEvents()

	// Open Shop
	shop.Open()

	// Send Random customers at random intervals
	wg := &sync.WaitGroup{}
	wg.Add(DAILY_CUSTOMERS)
	go sendRandomCustomer(shop, wg, seatsMutex)

	wg.Wait()

	// After last customer is served wait for a bit...
	time.Sleep(5 * time.Second)

	// Close Shop
	closeShopChan <- struct{}{}
	if !shop.isOpen {
		fmt.Println(std.PrintC(std.Purple, "-----------------------------------------------------------------------------------"))
		fmt.Println(std.PrintC(std.Purple, "SHOP CLOSED"))
		fmt.Println(std.PrintC(std.Purple, "-----------------------------------------------------------------------------------"))
		fmt.Println("Jobs Completed:", len(shop.jobs.completed))
	}

}

// Goroutine that sends customers at random intervals
func sendRandomCustomer(shop *BarberShop, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 1; i <= DAILY_CUSTOMERS; i++ {
		defer wg.Done()

		// Sleep for a random interval (Gap between each customer)
		time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)

		// Handle Entry of customer
		shop.handleCustomerEntry(&Customer{customerNumber: i, requirement: generateRandomCustomerRequirement()}, mutex)
	}
}

// Generate random requirement
func generateRandomCustomerRequirement() string {
	return SHOP_SERVICES[rand.Intn(len(SHOP_SERVICES)-1)]
}
