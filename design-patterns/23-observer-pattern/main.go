package main

import (
	"fmt"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/23-observer-pattern/lib"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

/****************************************************************************/
// (1) Basic Observer pattern implementation
/****************************************************************************/
// Observer defines the interface for the observer.
type Observer interface {
	Update(string)
}

// Observable represents the subject being observed.
type Observable struct {
	observers []Observer
}

// Register attaches an observer to the observable.
func (s *Observable) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Deregister detaches an observer from the observable.
func (s *Observable) Deregister(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

// Notify notifies all observers of a state change.
func (s *Observable) Notify(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserver represents a concrete observer.
type ConcreteObserver struct {
	Name string
}

// Update implements the Observer interface.
func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s received message: %s\n", o.Name, message)
}

/****************************************************************************/
// (2) Message Broker implementation
/****************************************************************************/
type Person struct {
	Name string
	Age  int
}

type Book struct {
	Name, Author string
}

type Stream[T interface{}] struct {
	Data []T
}

// A person data set
var personData = []Person{
	{"Jayanta Samaddar", 32},
	{"Rohit Saha", 33},
	{"Jaydeep Samaddar", 25},
	{"Raj Shaw", 31},
	{"Neelapravo Dasgupta", 30},
	{"Bhargav Sidagam", 32},
	{"Anil Biswakarma", 27},
	{"Harshvardhan Bhan", 34},
}

// A book data set
var bookData = []Book{
	{"Alice in Wonderland", "Lewis Caroll"},
	{"Paradise Lost", "John Milton"},
	{"Inferno", "Dante"},
	{"Illiad", "Homer"},
	{"Harry Potter and the Goblet of Fire", "J.K. Rowling"},
	{"You Can Win", "Shiv Khera"},
	{"Wings of Fire", "Dr. A.P.J. Abdul Kalam"},
	{"Lord of the Rings", "J.R.R. Tolkien"},
}

// A second book data set
var bookData2 = []Book{
	{"Angels and Demons", "Dan Brown"},
	{"How to Win Friends and Influence People", "Dale Carnegie"},
	{"Eragon", "Christopher Paolini"},
	{"Odyssey", "Homer"},
	{"Harry Potter and the Half-Blood Prince", "J.K. Rowling"},
	{"You Can Sell", "Shiv Khera"},
	{"India 2020", "Dr. A.P.J. Abdul Kalam"},
	{"Ramayana", "Valmiki"},
}

var doneCh = make(chan struct{})

/****************************************************************************/
// Main Function
/****************************************************************************/
func main() {
	/****************************************************************************/
	// (1) Demo: Basic Observer pattern implementation
	/****************************************************************************/
	observable := &Observable{}

	// Create observers
	observer1 := &ConcreteObserver{Name: "Observer 1"}
	observer2 := &ConcreteObserver{Name: "Observer 2"}

	// Register observers with the observable
	observable.Register(observer1)
	observable.Register(observer2)

	// Notify observers
	observable.Notify("Hello World")

	// Deregister observer1
	observable.Deregister(observer1)

	// Notify observers again
	observable.Notify("Goodbye")

	/****************************************************************************/
	// (1) Demo: Message Queue (Implementation of Mediator-Observer Pattern)
	/****************************************************************************/
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	fmt.Println(std.Purple + "\t(2) Demo: Message Queue (Implementation of Mediator-Observer Pattern)" + std.Reset)
	fmt.Println(std.Purple + "-------------------------------------------------------------------------------------------------" + std.Reset)
	topicName := "person-topic"
	topicName2 := "book-topic"
	msgQ := lib.NewMessageQueue("events")

	c := lib.NewConsumer().SetID("consumer-1").BindQueue(msgQ)
	c.Subscribe(topicName)
	c.Subscribe(topicName2)
	c.Subscribe("non-existent-topic")

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for {
	// 		select {
	// 		case qNotification := <-msgQ.GetQueueEvents():
	// 			fmt.Println("Received event notification")
	// 			fmt.Println("Queue event notification:", qNotification)
	// 		case <-doneCh:
	// 			return
	// 		}
	// 	}
	// }()

	fmt.Println("Subscribed Topics:", c.GetTopics()) // Subscribed Topics: [person-topic book-topic non-existent-topic]
	c.Unsubscribe("non-existent-topic")
	fmt.Println("Subscribed Topics after unsubscribing:", c.GetTopics()) // Subscribed Topics after unsubscribing: [person-topic book-topic]

	// Consume events for one subscribed topic
	// go c.Consume(topicName, func(event lib.Event) {
	// 	msg := fmt.Sprintf("Event being consumed by Consumer %s:", c.ID())
	// 	fmt.Println(std.Green+msg, event.String(), std.Reset)
	// }, doneCh)

	// Consume events for all subscribed topics
	go c.ConsumeAll(func(event lib.Event) {
		msg := fmt.Sprintf("Event being consumed by Consumer %s:", c.ID())
		fmt.Println(std.Green+msg, event.String(), std.Reset)
	}, doneCh)

	fmt.Println(std.Cyan + "-------------------------------------------------------------------------------------------------" + std.Reset)
	// Publish Events to the Queue under a topic ("person-topic")
	personDataStream := Stream[Person]{personData}
	publishEvents(simulateEventGeneration(msgQ, topicName, personDataStream))

	fmt.Println(std.Cyan + "-------------------------------------------------------------------------------------------------" + std.Reset)
	// Publish more events to the queue under a different topic ("book-topic")
	bookDataStream := Stream[Book]{bookData}
	publishEvents(simulateEventGeneration(msgQ, topicName2, bookDataStream))

	fmt.Println(std.Cyan + "-------------------------------------------------------------------------------------------------" + std.Reset)
	// Publish more events to the queue under the "book-topic", however unsubscribe consumer first
	c.Unsubscribe(topicName2)
	bookDataStream2 := Stream[Book]{bookData2}
	publishEvents(simulateEventGeneration(msgQ, topicName2, bookDataStream2))
}

// Publish events to the Queue
func publishEvents(eventGenerator chan lib.Observable) {
	for event := range eventGenerator {
		fmt.Println(std.Yellow+"Event being published:", event, std.Reset)
		event.Publish()
	}
	time.Sleep(100 * time.Millisecond)
}

// Simulate a generation of data stream
func simulateEventGeneration[T interface{}](queue *lib.MessageQueue, topic string, stream Stream[T]) chan lib.Observable {
	producerCh := make(chan lib.Observable)
	go func() {
		defer close(producerCh)
		for _, p := range stream.Data {
			producerCh <- lib.NewEvent(queue, topic, p)
		}
	}()
	return producerCh
}
