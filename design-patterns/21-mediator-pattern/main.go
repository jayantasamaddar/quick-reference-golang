package main

import "fmt"

/*****************************************************************************/
// Participants that want to communicate with each other
/*****************************************************************************/
type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

// Factory Function to create a new person
func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Send(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) SendPrivate(receiver, message string) {
	p.Room.Message(p.Name, receiver, message)
}

/*****************************************************************************/
// ChatRoom: This is the Mediator
/*****************************************************************************/

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(sender, message string) {
	for _, p := range c.people {
		if sender != p.Name {
			p.Receive(sender, message)
		}
	}
}

func (c *ChatRoom) Message(sender, receiver, message string) {
	for _, p := range c.people {
		if receiver == p.Name {
			p.Receive(sender, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	c.Broadcast("Room", fmt.Sprintf("%s has joined the chat!", p.Name))
	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	// Chat Room
	room := ChatRoom{}

	// Users
	jayanta := NewPerson("Jayanta")
	rohit := NewPerson("Rohit")

	// Join Room
	room.Join(jayanta)
	room.Join(rohit)

	jayanta.Send("Hi Room!")
	rohit.Send("Oh, Hi Jayanta!")

	// New Person joins!
	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Send("Hi Everyone!")

	// Private Message
	jayanta.SendPrivate(simon.Name, "Glad you could join us!")

	/*
		[Jayanta's chat session]: Room: Rohit has joined the chat!
		[Rohit's chat session]: Jayanta: Hi Room!
		[Jayanta's chat session]: Rohit: Oh, Hi Jayanta!
		[Jayanta's chat session]: Room: Simon has joined the chat!
		[Rohit's chat session]: Room: Simon has joined the chat!
		[Jayanta's chat session]: Simon: Hi Everyone!
		[Rohit's chat session]: Simon: Hi Everyone!
		[Simon's chat session]: Jayanta: Glad you could join us!
	*/

}
