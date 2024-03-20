# Table of Contents

- [Table of Contents](#table-of-contents)
- [Observer Pattern](#observer-pattern)
  - [Introduction](#introduction)
  - [Use Cases](#use-cases)
  - [Implementation in Go](#implementation-in-go)
    - [Basic Implementation](#basic-implementation)
  - [Observer-Mediator Pattern](#observer-mediator-pattern)
    - [The Message Queue](#the-message-queue)
    - [Implementation of the Message Queue](#implementation-of-the-message-queue)
- [Summary](#summary)

---

# Observer Pattern

## Introduction

The Observer Pattern is a behavioral design pattern where an object, known as the **Subject** or **Observable**, maintains a list of its dependents, called **Observers**, and notifies them of any changes in state, typically by calling one of their methods. This pattern is useful in scenarios where multiple objects need to be notified of changes in another object's state without coupling the subjects to the observers.

> An **Observer** is an object that wishes to be informed about events happening in the system. The entity generating the events is an **Observable**.

---

## Use Cases

The Observer pattern is widely used in software development and can be found in various libraries and frameworks across different programming languages. Some practical use cases of the Observer pattern include:

1. **Graphical User Interfaces (GUIs)**: In GUI frameworks like Swing in Java or GTK+ in Python, the Observer pattern is used extensively. For example, a button click event can notify multiple listeners (observers) such as event handlers or components interested in that event.

2. **Model-View-Controller (MVC) Architecture**: The Observer pattern is fundamental to the MVC architecture. In MVC, the model notifies the views of any changes in its state, and the views update themselves accordingly. This decouples the model from the views, enabling better maintainability and scalability.

3. **Event-driven Programming**: In event-driven programming paradigms, such as in web development with JavaScript or Node.js, the Observer pattern is employed. Event emitters (subjects) emit events, and various event listeners (observers) subscribe to these events and react accordingly.

4. **Databases and ORM**: Object-Relational Mapping (ORM) frameworks often use the Observer pattern. For instance, an ORM might notify interested parties whenever a database record is updated, allowing other parts of the application to react to those changes.

5. **Publish-Subscribe Systems**: Systems that implement publish-subscribe messaging patterns use the Observer pattern under the hood. Publishers (subjects) publish messages, and subscribers (observers) receive those messages based on their subscriptions.

**Popular libraries and frameworks that use the Observer pattern include**:

1. **Java's java.util.Observer/Observable**: Java provides built-in support for the Observer pattern through the `java.util.Observer` interface and the `java.util.Observable` class.

2. **AngularJS and Angular (JavaScript/TypeScript)**: These frontend frameworks employ the Observer pattern extensively for data binding and event handling. Components can subscribe to changes in data models and update the UI accordingly.

3. **React (JavaScript)**: Although React uses its own concepts like props and state, it can be considered to follow a unidirectional data flow architecture, which shares similarities with the Observer pattern. Components subscribe to changes in state and re-render themselves when necessary.

4. **.NET's INotifyPropertyChanged Interface**: In .NET, the `INotifyPropertyChanged` interface is used to implement the Observer pattern, allowing objects to notify observers of changes in their properties.

5. **Django (Python)**: Django's signals framework enables decoupled applications to get notified when certain actions occur elsewhere in the application. This is achieved using the Observer pattern.

---

## Implementation in Go

### Basic Implementation

In Go, we can demonstrate the Observer Pattern using interfaces and channels.

Here's an example implementation:

```go
package main

import "fmt"

// Observer defines the interface for the observer.
type Observer interface {
    Update(string)
}

// Subject represents the subject being observed.
type Subject struct {
    observers []Observer
}

// Register attaches an observer to the subject.
func (s *Subject) Register(observer Observer) {
    s.observers = append(s.observers, observer)
}

// Deregister detaches an observer from the subject.
func (s *Subject) Deregister(observer Observer) {
    for i, obs := range s.observers {
        if obs == observer {
            s.observers = append(s.observers[:i], s.observers[i+1:]...)
            return
        }
    }
}

// Notify notifies all observers of a state change.
func (s *Subject) Notify(message string) {
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

func main() {
    subject := &Subject{}

    // Create observers
    observer1 := &ConcreteObserver{Name: "Observer 1"}
    observer2 := &ConcreteObserver{Name: "Observer 2"}

    // Register observers with the subject
    subject.Register(observer1)
    subject.Register(observer2)

    // Notify observers
    subject.Notify("Hello World")

    // Deregister observer1
    subject.Deregister(observer1)

    // Notify observers again
    subject.Notify("Goodbye")
}
```

---

## Observer-Mediator Pattern

### The Message Queue

In certain implementations, a Message Queue can be seen as a combination of the Observer pattern and the Mediator pattern:

- **Observer Pattern**: In a messaging queue system, publishers (or producers) send messages to the queue, while subscribers (or consumers) receive and process those messages. This behavior resembles the Observer pattern, where the queue acts as the subject (or publisher) that notifies interested parties (subscribers or consumers) of new messages as they arrive.

- **Mediator Pattern**: The messaging queue acts as a mediator between publishers and subscribers. Instead of publishers directly communicating with subscribers, they interact with the queue, which then distributes messages to the appropriate subscribers. This decouples publishers from subscribers, promoting loose coupling and scalability. The queue serves as a central mediator, coordinating communication between different components of the system without them needing to be aware of each other.

In summary, while a messaging queue primarily functions as a mediator for communication between different parts of a system, its interaction model with publishers and subscribers resembles that of the Observer pattern. Therefore, it can be considered a combination of both patterns, effectively managing communication and decoupling components in distributed systems.

---

### Implementation of the Message Queue

Find the complete implementation of the `MessageQueue` system at **[`lib/MessageQueue.go`](./lib/MessageQueue.go)**:

In **`main.go`**:

```go
package main

import (
	"fmt"
	"time"

	"github.com/jayantasamaddar/quick-reference-golang/design-patterns/23-observer-pattern/lib"
	"github.com/jayantasamaddar/quick-reference-golang/golang/std"
)

/****************************************************************************/
// Message Queue implementation
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

var doneCh = make(chan struct{})

/****************************************************************************/
// Main Function
/****************************************************************************/
func main() {
	/****************************************************************************/
	// Demo: Message Queue (Implementation of Mediator-Observer Pattern)
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

	// Publish Events to the Queue
	personDataStream := Stream[Person]{personData}
	publishEvents(simulateEventGeneration(msgQ, topicName, personDataStream))

	// Publish more events to the queue under a different topic
	bookDataStream := Stream[Book]{bookData}
	publishEvents(simulateEventGeneration(msgQ, topicName2, bookDataStream))

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
```

In this implementation, the Mediator pattern and Observer pattern are used to facilitate communication between components in a decoupled manner.

1. **Mediator Pattern**:

   - The `MessageQueue` struct acts as the Mediator. It centralizes communication between different components (`Topic` and `Consumer`).
   - The `MessageQueue` maintains a collection of topics and handles the routing of events between producers (publishers) and consumers.
   - Topics are created and managed by the `MessageQueue`, which allows producers and consumers to interact without needing to know about each other directly.
   - When an event is published, the `MessageQueue` routes it to the appropriate topic, ensuring that consumers subscribed to that topic receive the event.

2. **Observer Pattern**:

   - Consumers act as Observers. They subscribe to specific topics they are interested in consuming events from.
   - Consumers do not need to know about each other or about the producers. They simply subscribe to topics they care about.
   - When an event is published to a topic, all subscribed consumers receive the event asynchronously via channels.
   - The Observer pattern decouples producers from consumers, allowing for more flexible and scalable systems.

3. **Usage in the Code**:

   - The `MessageQueue` struct encapsulates the logic for managing topics and routing events.
   - Topics manage the list of consumers interested in events and handle the distribution of events to those consumers.
   - Consumers subscribe to topics they are interested in and consume events asynchronously.
   - Producers (Observable) publish events to the `MessageQueue`, which then routes them to the appropriate topics for consumption.

4. **Other Design Patterns**:
   - The code also uses the Factory pattern implicitly when creating instances of events, topics, and consumers. For example, `NewEvent`, `NewTopic`, and `NewConsumer` functions are responsible for creating instances of their respective structs.
   - The code also demonstrates the usage of channels for asynchronous communication, which is a common pattern in Go for concurrent programming.

Overall, by employing the Mediator and Observer patterns, the code achieves a flexible and scalable architecture where components are loosely coupled, making it easier to maintain and extend the system.

---

# Summary

---
